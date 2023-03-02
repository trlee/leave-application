package handlers

import (
	"log"
	"net/http"
	"strconv"
)

type payloadUpdateReason struct {
	ID     string `json:"ID"`
	Reason string `json:"updateReason"`
}

func (rep *Repository) CancelLeaveApplication(w http.ResponseWriter, r *http.Request) {
	var p payloadUpdateReason
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	id, reason := convertPayloadUpdateReason(p)
	result, err := rep.App.Models.LeaveApplication.CancelLeaveApplication(id, reason)
	if err != nil {
		log.Println("error while update: ", err)
	}

	log.Println("leave application successfully updated: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave application successfully updated")

}

func (rep *Repository) ApproveLeaveApplication(w http.ResponseWriter, r *http.Request) {
	var p payloadUpdateReason
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	id, reason := convertPayloadUpdateReason(p)
	result, err := rep.App.Models.LeaveApplication.ApproveLeaveApplication(id, reason)
	if err != nil {
		log.Println("error while update: ", err)
	}

	log.Println("leave application successfully updated: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave application successfully updated")

}

func (rep *Repository) RejectLeaveApplication(w http.ResponseWriter, r *http.Request) {
	var p payloadUpdateReason
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	id, reason := convertPayloadUpdateReason(p)
	result, err := rep.App.Models.LeaveApplication.RejectLeaveApplication(id, reason)
	if err != nil {
		log.Println("error while update: ", err)
	}

	log.Println("leave application successfully updated: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave application successfully updated")

}

func convertPayloadUpdateReason(p payloadUpdateReason) (int, string) {
	id, _ := strconv.Atoi(p.ID)
	reason := p.Reason

	return id, reason
}
