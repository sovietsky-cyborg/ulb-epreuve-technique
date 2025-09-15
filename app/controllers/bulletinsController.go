package controllers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"ucl-epreuve-technique/app/models"
	"ucl-epreuve-technique/app/utils"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Bulletin struct {
	Nom              string            `json:"nom"`
	Prenom           string            `json:"prenom"`
	Matricule        string            `json:"matricule"`
	Annee            int               `json:"annee"`
	TotalCredits     int               `json:"total_credits"`
	ValidatedCredits int               `json:"validated_credits"`
	WeightedAverage  int               `json:"weighted_average"`
	IsSuccess        bool              `json:"is_success"`
	ListeCours       []ListeCoursNotes `json:"liste_cours"`
}

type ListeCoursNotes struct {
	models.ListeCours
	Note int `json:"note"`
}

// sort Interface is implemented to support Sorting by Mnemonique on ListeCoursNotes
type ByMnemonique []ListeCoursNotes

func (m ByMnemonique) Len() int           { return len(m) }
func (m ByMnemonique) Less(i, j int) bool { return m[i].Mnemonique < m[j].Mnemonique }
func (m ByMnemonique) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

var GetBulletinHandler = func(w http.ResponseWriter, r *http.Request) interface{} {

	var bulletin Bulletin

	vars := mux.Vars(r)
	annee, _ := strconv.Atoi(vars["annee"])
	matricule := vars["matricule"]

	validate := validator.New()

	err := validate.Var(matricule, "required,alphanum")
	if err != nil {
		return utils.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	err = validate.Var(annee, "gte=0,lte=3,number")
	if err != nil {
		return utils.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	client := utils.GetClient()

	body, err := client.GetData("/inscriptions?matricule=" + matricule)
	if err != nil {
		return err
	}

	var result []models.ListeInscriptions
	var inscription models.ListeInscriptions
	err = json.Unmarshal(body, &result)
	inscription = result[0]

	bulletin.Nom = inscription.Nom
	bulletin.Prenom = inscription.Prenom
	bulletin.Matricule = inscription.Matricule
	bulletin.Annee = inscription.AnneeEtude

	// Just Get all the note for this inscription,
	// no need to wait O(N_Courses * API_Calls) by calling swagger endpoint for each mnemonique
	body, err = client.GetData("/notes?matricule=" + inscription.Matricule)
	var notes []models.ListeNotes
	err = json.Unmarshal(body, &notes)

	notesByCours := make(map[string]models.ListeNotes)
	// instead of that, we keep track of notes by mnemonique
	for _, note := range notes {
		notesByCours[note.Mnemonique] = note
	}
	has10OnAllCourses := true

	for _, cour := range inscription.CoursJson {

		body, err = client.GetData("/cours?mnemonique=" + cour)
		var currentCours []models.ListeCours
		err = json.Unmarshal(body, &currentCours)
		// then, recover them when iterating throughout currentInscription.CoursJson
		currentNote := notesByCours[currentCours[0].Mnemonique].Note
		bulletin.TotalCredits += currentCours[0].Credit
		bulletin.ListeCours = append(bulletin.ListeCours, ListeCoursNotes{
			ListeCours: currentCours[0],
			Note:       currentNote,
		})

		if currentNote < 10 {
			has10OnAllCourses = false
		} else {
			bulletin.ValidatedCredits += currentCours[0].Credit
		}

		bulletin.WeightedAverage += currentCours[0].Credit * notesByCours[currentCours[0].Mnemonique].Note

	}
	bulletin.WeightedAverage /= bulletin.TotalCredits

	if (has10OnAllCourses && bulletin.WeightedAverage > 10) || bulletin.ValidatedCredits > 60 {
		bulletin.IsSuccess = true
	}
	sort.Sort(ByMnemonique(bulletin.ListeCours))

	return bulletin
}
