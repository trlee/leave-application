package handlers

import (
	"leave/pkg/pg"
	"log"
	"net/http"
	"strconv"
)

type payloadNewApplication struct {
	Email     string `json:"email"`
	LeaveType string `json:"leaveType"`
	LeaveDate string `json:"leaveDate"`
	HalfDay   string `json:halfDay`
	Reason    string `json:"applyReason"`
}

func (rep *Repository) ApplyLeave(w http.ResponseWriter, r *http.Request) {
	var p payloadNewApplication
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	newLeaveApplication := convertPayloadNewApplication(p)
	result, err := rep.App.Models.LeaveApplicationTest.ApplyLeave(newLeaveApplication)
	if err != nil {
		log.Println("error while inserted: ", err)
	}

	log.Println("leave application successfully inserted: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave application successfully inserted")

}

func convertPayloadNewApplication(p payloadNewApplication) pg.LeaveApplicationTest {
	leaveType, _ := strconv.Atoi(p.LeaveType)
	halfDay, _ := strconv.ParseBool(p.HalfDay)

	return pg.LeaveApplicationTest{
		Email:     p.Email,
		LeaveType: leaveType,
		LeaveDate: p.LeaveDate,
		HalfDay:   halfDay,
		Reason:    p.Reason,
	}

}
