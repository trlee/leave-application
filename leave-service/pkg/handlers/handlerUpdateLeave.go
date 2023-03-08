package handlers

import (
	"leave/pkg/pg"
	"log"
	"net/http"
	"strconv"
)

type payloadUpdateLeaveType struct {
	ID                     string `json:"ID"`
	LeaveType              string `json:"name"`
	Limit                  string `json:"maxLimit"`
	EntitlementCalculation string `json:"entitlementCalc"`
	Gender                 string `json:"gender"`
	Unpaid                 string `json:"isUnpaid"`
	AttachmentMandatory    string `json:"isAttachmentMandatory"`
	EncashmentLeave        string `json:"isEncashmentLeave"`
}

func (rep *Repository) UpdateLeave(w http.ResponseWriter, r *http.Request) {
	var p payloadUpdateLeaveType
	err := rep.readJSON(w, r, &p)
	if err != nil {
		rep.errorJSON(w, err)
		return
	}
	updatedLeaveType := convertPayloadUpdatedLeaveType(p)
	result, err := rep.App.Models.LeaveType.UpdateLeave(updatedLeaveType)
	if err != nil {
		log.Println("error while inserted: ", err)
	}
	log.Println("leave type successfully inserted: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave type successfully inserted")
}

func convertPayloadUpdatedLeaveType(p payloadUpdateLeaveType) pg.LeaveType {
	log.Println("ECalc", p.EntitlementCalculation)
	id, _ := strconv.Atoi(p.ID)
	limit, _ := strconv.Atoi(p.Limit)
	gender, _ := strconv.Atoi(p.Gender)
	entitlementCalc, err := strconv.ParseFloat(p.EntitlementCalculation, 32)
	if err != nil {
		log.Println("ERROR! ", err)
	}
	unpaid, _ := strconv.ParseBool(p.Unpaid)
	attachmentMand, _ := strconv.ParseBool(p.AttachmentMandatory)
	encashmentLeave, _ := strconv.ParseBool(p.EncashmentLeave)

	return pg.LeaveType{
		ID:                     id,
		LeaveType:              p.LeaveType,
		Unpaid:                 unpaid,
		Limit:                  limit,
		EntitlementCalculation: float32(entitlementCalc),
		Gender:                 gender,
		AttachmentMandatory:    attachmentMand,
		EncashmentLeave:        encashmentLeave,
	}

}
