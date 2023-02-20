package handlers

import (
	"encoding/json"
	"leave/pkg/pg"
	"log"
	"net/http"
)

func (rep *Repository) LeaveReport(w http.ResponseWriter, r *http.Request) {
	myc := r.Context().Value(httpContext).(httpContextStruct)
	requestPayload := ReportPayload{
		Email: myc.User.Email,
	}
	all, err := rep.App.Models.EmployeeLeave.GetAllLeaves(requestPayload.Email)
	if err != nil {
		log.Println("error while fetching all leaves", err)
	}

	type myJsonResp struct {
		Data []*pg.EmployeeLeave
	}
	resp := myJsonResp{
		Data: all,
	}

	json.NewEncoder(w).Encode(&resp)
}
