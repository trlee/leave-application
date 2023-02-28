package main

import (
	"leave/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/leave", handlers.Repo.LeaveReport)
	mux.HandleFunc("/leaveApply", handlers.Repo.ApplyLeave)
	mux.HandleFunc("/leaveTypeAdd", handlers.Repo.AddLeave)
	mux.HandleFunc("/getAllLeaveApplication", handlers.Repo.GetAllLeaveApplication)
	mux.HandleFunc("/getAllLeaveType", handlers.Repo.GetAllLeaveTypes)

	return mux
}
