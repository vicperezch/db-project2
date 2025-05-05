package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"net/http"
)

func GetLoans(w http.ResponseWriter, r *http.Request) {
	var loans []model.Loan

	result := database.DB.Joins("Client").Find(&loans)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get loans.", http.StatusInternalServerError)
		return
	}

	util.RespondWithJSON(w, loans)
}
