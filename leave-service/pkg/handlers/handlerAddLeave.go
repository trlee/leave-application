package handlers

import (
	"leave/pkg/pg"
	"log"
	"net/http"
	"strconv"
)

type payloadLeaveType struct {
	LeaveType              string `json:"name"`
	Limit                  string `json:"maxLimit"`
	EntitlementCalculation string `json:"entitlementCalculation"`
	Gender                 string `json:"gender"`
	Unpaid                 string `json:"isUnpaid"`
	AttachmentMandatory    string `json:"isAttachmentMandatory"`
	EncashmentLeave        string `json:"isEncashmentLeave"`
}

func (rep *Repository) AddLeave(w http.ResponseWriter, r *http.Request) {
	var p payloadLeaveType
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	newLeaveType := convertPayloadNewLeaveType(p)
	result, err := rep.App.Models.LeaveType.AddLeave(newLeaveType)
	if err != nil {
		log.Println("error while inserted: ", err)
	}

	log.Println("leave type successfully inserted: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave type successfully inserted")

}

func convertPayloadNewLeaveType(p payloadLeaveType) pg.LeaveType {
	limit, _ := strconv.Atoi(p.Limit)
	gender, _ := strconv.Atoi(p.Gender)
	entitlementCalc, _ := strconv.ParseFloat(p.EntitlementCalculation, 32)
	unpaid, _ := strconv.ParseBool(p.Unpaid)
	attachmentMand, _ := strconv.ParseBool(p.AttachmentMandatory)
	encashmentLeave, _ := strconv.ParseBool(p.EncashmentLeave)

	return pg.LeaveType{
		ID:                     0,
		LeaveType:              p.LeaveType,
		Unpaid:                 unpaid,
		Limit:                  limit,
		EntitlementCalculation: float32(entitlementCalc),
		Gender:                 gender,
		AttachmentMandatory:    attachmentMand,
		EncashmentLeave:        encashmentLeave,
	}

}
