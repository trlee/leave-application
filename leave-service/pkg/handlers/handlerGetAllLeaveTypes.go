package handlers

import (
	"log"
	"net/http"
)

// API to fetch all leave types
func (rep *Repository) GetAllLeaveType(w http.ResponseWriter, r *http.Request) {

	// Get all leave types form database
	all, err := rep.App.Models.LeaveType.GetAllLeaveTypes()

	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	log.Println(all)

	// Return answer to front-end
	answer := jsonResponse{
		Error:   false,
		Message: "All Leave Types collected",
		Data:    all,
	}

	log.Println(all)

	rep.writeJSON(w, http.StatusAccepted, answer)
}
