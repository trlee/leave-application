package main

import (
	"leave/pkg/handlers"
	"net/http"
)

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// Statics route public
	mux.HandleFunc("/leave", handlers.Repo.LeaveReport)
	mux.HandleFunc("/leave/apply", handlers.Repo.ApplyLeave)

	return mux
}
