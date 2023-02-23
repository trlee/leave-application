package handlers

import "net/http"

// API to fetch all leave application
func (rep *Repository) GetAllLeaveApplication(w http.ResponseWriter, r *http.Request) {
	// Get all leave application from database
	email := "admin@thundersoft.com"
	all, err := rep.App.Models.LeaveApplication.GetAllApplications(email)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}

	// Return answer to front-end
	answer := jsonResponse{
		Error:   false,
		Message: "All Leave Applications collected",
		Data:    all,
	}

	rep.writeJSON(w, http.StatusAccepted, answer)
}
