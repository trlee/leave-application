package main

import (
	"leave/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/leaveApply", handlers.Repo.ApplyLeave)
	mux.HandleFunc("/leaveTypeAdd", handlers.Repo.AddLeave)
	mux.HandleFunc("/getAllLeaveApplication", handlers.Repo.GetAllLeaveApplication)
	mux.HandleFunc("/getAllLeaveType", handlers.Repo.GetAllLeaveType)
	mux.HandleFunc("/getAllLeaveReport", handlers.Repo.LeaveReport)
	mux.HandleFunc("/cancelLeaveApplication", handlers.Repo.CancelLeaveApplication)
	mux.HandleFunc("/approveLeaveApplication", handlers.Repo.ApproveLeaveApplication)
	mux.HandleFunc("/rejectLeaveApplication", handlers.Repo.RejectLeaveApplication)
	mux.HandleFunc("/leaveTypeUpdate", handlers.Repo.UpdateLeave)
	return mux
}
