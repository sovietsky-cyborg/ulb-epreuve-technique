package controllers

import (
	"errors"
	"net/http"
	"ucl-epreuve-technique/app/models"
	"ucl-epreuve-technique/app/utils"
)

var GetNotesHandler = func(w http.ResponseWriter, r *http.Request) interface{} {

	notes, err := models.GetNotes()

	if len(notes) == 0 {
		return utils.StatusError{Code: http.StatusNotFound, Err: errors.New("no notes available")}
	}
	if err != nil {
		return utils.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return notes

}
