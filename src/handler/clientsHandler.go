package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"net/http"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	var clients []model.Client

	result := database.DB.Find(&clients)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get clients.", http.StatusInternalServerError)
		return
	}

	util.RespondWithJSON(w, clients)
}
