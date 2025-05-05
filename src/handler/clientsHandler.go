package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"encoding/csv"
	"net/http"
	"strconv"
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

func ExportClientsToCSV(w http.ResponseWriter, r *http.Request) {
	var clients []model.Client

	result := database.DB.Find(&clients)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get clients.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment;filename=clients.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Last name", "Address", "Phone", "Email"})

	for _, c := range clients {
		writer.Write([]string{
			strconv.FormatUint(uint64(c.ID), 10),
			c.Names,
			c.LastNames,
			c.Address,
			c.Phone,
			c.Email,
		})
	}
}
