package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"net/http"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []model.Author

	result := database.DB.Find(&authors)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get clients.", http.StatusInternalServerError)
		return
	}

	util.RespondWithJSON(w, authors)
}
