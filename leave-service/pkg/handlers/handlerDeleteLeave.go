package handlers

import (
	"log"
	"net/http"
)

func (rep *Repository) DeleteLeave(w http.ResponseWriter, r *http.Request) {

	var p any
	err := rep.readJSON(w, r, &p)
	if err != nil {
		log.Println(err)
		return
	}

	result, err := rep.App.Models.LeaveType.DeleteLeave(p)
	if err != nil {
		log.Println("error while deleting: ", err)
	}

	log.Println("leave type successfully deleted: ", result)
	rep.writeJSON(w, http.StatusAccepted, "leave type successfully deleted")

}
