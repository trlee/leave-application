package handlers

import (
	"authentication/pkg/pg"
	"fmt"
	"net/http"
	"strconv"
)

func (rep *Repository) Approve(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	i, err := strconv.Atoi(r.FormValue("duration_input"))
	if err != nil {
		fmt.Println(err)
	}
	requestPayload := LeavePayload{
		Email:    r.FormValue("email_input"),
		Duration: i,
		Reason:   r.FormValue("reason_input"),
	}

	fmt.Println("I'm here!", requestPayload.Email, requestPayload.Duration, requestPayload.Reason)

	leave, err := rep.App.Models.Leave.ApproveLeave(pg.Leave(requestPayload))
	if err != nil {
		fmt.Println(leave, err)
	} else {
		fmt.Println(leave)
		http.Redirect(w, r, "http://localhost/home", http.StatusSeeOther)
	}

}
