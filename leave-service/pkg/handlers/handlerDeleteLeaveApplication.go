package handlers

import (
	"log"
	"net/http"
)

func (rep *Repository) DeleteLeaveApplication(w http.ResponseWriter, r *http.Request) {

	var p any
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println(err)
		return
	}

	result, err := rep.App.Models.LeaveApplication.DeleteLeaveApplication(p)
	if err != nil {
		log.Println("error while deleting: ", err)
	}

	log.Println("leave type successfully deleted: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave type successfully deleted")

}
