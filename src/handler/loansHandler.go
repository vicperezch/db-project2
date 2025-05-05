package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"encoding/csv"
	"net/http"
	"strconv"
)

func GetLoans(w http.ResponseWriter, r *http.Request) {
	var loans []model.Loan

	result := database.DB.Find(&loans)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get loans.", http.StatusInternalServerError)
		return
	}

	util.RespondWithJSON(w, loans)
}

func ExportLoansToCSV(w http.ResponseWriter, r *http.Request) {
	var loans []model.Loan

	result := database.DB.Find(&loans)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get loans.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment;filename=loans.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"ID", "ClientID", "EmployeeID", "EditionID", "Loan date", "Return date"})

	for _, c := range loans {
		writer.Write([]string{
			strconv.FormatUint(uint64(c.ID), 10),
			strconv.FormatUint(uint64(c.ClientID), 10),
			strconv.FormatUint(uint64(c.EmployeeID), 10),
			strconv.FormatUint(uint64(c.EditionID), 10),
			c.LoanDate.String(),
			c.ReturnDate.String(),
		})
	}
}
