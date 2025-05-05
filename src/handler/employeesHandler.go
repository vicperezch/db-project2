package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"encoding/csv"
	"net/http"
	"strconv"
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

func ExportEmployeesToCSV(w http.ResponseWriter, r *http.Request) {
	var employees []model.Employee

	result := database.DB.Find(&employees)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get employees.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment;filename=employees.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name", "Last name", "Hiring date"})

	for _, c := range employees {
		writer.Write([]string{
			strconv.FormatUint(uint64(c.ID), 10),
			c.Names,
			c.LastNames,
			c.HiringDate.String(),
		})
	}
}
