package handlers

import (
	"encoding/json"
	"leave/pkg/pg"
	"log"
	"net/http"
)

func (rep *Repository) LeaveApplication(w http.ResponseWriter, r *http.Request) {
	myc := r.Context().Value(httpContext).(httpContextStruct)
	all, err := rep.App.Models.LeaveApplication.GetAllApplications(myc.User.Email)
	if err != nil {
		log.Println("error while fetching all applications", err)
	}

	type myJsonResp struct {
		Data []*pg.LeaveApplication
	}
	resp := myJsonResp{
		Data: all,
	}

	json.NewEncoder(w).Encode(&resp)
}
