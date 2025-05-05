package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"encoding/csv"
	"net/http"
	"strconv"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []model.Author

	result := database.DB.Find(&authors)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get authors.", http.StatusInternalServerError)
		return
	}

	util.RespondWithJSON(w, authors)
}

func ExportAuthorsToCSV(w http.ResponseWriter, r *http.Request) {
	var authors []model.Author

	result := database.DB.Find(&authors)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get authors.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment;filename=authors.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Last name", "Nationality"})

	for _, c := range authors {
		writer.Write([]string{
			strconv.FormatUint(uint64(c.ID), 10),
			c.Names,
			c.LastNames,
			c.Nationality,
		})
	}
}
