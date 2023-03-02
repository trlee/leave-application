package handlers

import (
	"log"
	"net/http"
)

func (rep *Repository) LeaveReport(w http.ResponseWriter, r *http.Request) {
	log.Println("I'm here!!")
	//myc := r.Context().Value(httpContext).(httpContextStruct)
	email := "admin@thundersoft.com"
	all, err := rep.App.Models.EmployeeLeave.GetAllLeaves(email)
	if err != nil {
		log.Println("error while fetching all leaves", err)
	}

	answer := jsonResponse{
		Error:   false,
		Message: "All Employee Leave",
		Data:    all,
	}

	log.Println(all)

	rep.writeJSON(w, http.StatusAccepted, answer)
}
