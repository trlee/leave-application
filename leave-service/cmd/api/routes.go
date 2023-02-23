package main

import (
	"leave/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/leave", handlers.Repo.LeaveReport)
	mux.HandleFunc("/leaveApplications", handlers.Repo.LeaveApplication)
	mux.HandleFunc("/leaveApply", handlers.Repo.ApplyLeave)
	mux.HandleFunc("/leaveAdd", handlers.Repo.AddLeave)
	mux.HandleFunc("/getAllLeaveApplication", handlers.Repo.GetAllLeaveApplication)

	return mux
}
