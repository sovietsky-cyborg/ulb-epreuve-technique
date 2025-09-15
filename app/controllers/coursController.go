package controllers

import (
	"errors"
	"net/http"
	"ucl-epreuve-technique/app/models"
	"ucl-epreuve-technique/app/utils"
)

var GetCoursHandler = func(w http.ResponseWriter, r *http.Request) interface{} {

	cours, err := models.GetCours()
	if len(cours) == 0 {
		return utils.StatusError{Code: http.StatusNotFound, Err: errors.New("no cours available")}
	}
	if err != nil {
		return utils.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return cours
}
