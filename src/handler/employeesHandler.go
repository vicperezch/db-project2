package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []model.Employee

	result := database.DB.Find(&employees)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get employees.", http.StatusInternalServerError)
		return
	}

	util.RespondWithJSON(w, employees)
}
