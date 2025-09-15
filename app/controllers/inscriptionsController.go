package controllers

import (
	"errors"
	"net/http"
	"ucl-epreuve-technique/app/models"
	"ucl-epreuve-technique/app/utils"
)

var GetInscriptionsHandler = func(w http.ResponseWriter, r *http.Request) interface{} {

	inscriptions, err := models.GetInscriptions()

	if len(inscriptions) == 0 {
		return utils.StatusError{Code: http.StatusNotFound, Err: errors.New("no inscription available")}
	}

	if err != nil {
		return utils.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return inscriptions

}
