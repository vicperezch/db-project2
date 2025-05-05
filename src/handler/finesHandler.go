package handler

import (
	"editorial-backend/database"
	"editorial-backend/model"
	"editorial-backend/util"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
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

func ExportFinesToCSV(w http.ResponseWriter, r *http.Request) {
	var fines []model.Fine

	result := database.DB.Find(&fines)
	if result.Error != nil {
		util.RespondWithError(w, "Could not get fines.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment;filename=fines.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	writer.Write([]string{"ID", "LoanID", "Reason", "Amount", "Payment date"})

	for _, c := range fines {
		var reason string = ""
		if c.Reason != nil {
			reason = *c.Reason
		}

		writer.Write([]string{
			strconv.FormatUint(uint64(c.ID), 10),
			strconv.FormatUint(uint64(c.LoanID), 10),
			reason,
			fmt.Sprintf("%.2f", c.Amount),
			c.PaymentDate.String(),
		})
	}
}
