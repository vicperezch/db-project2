package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"net/http"
)

func GetFines(w http.ResponseWriter, r *http.Request) {
	var fines []model.Fine

	result := database.DB.Find(&fines)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get fines.", http.StatusInternalServerError)
		return
	}

	util.RespondWithJSON(w, fines)
}
