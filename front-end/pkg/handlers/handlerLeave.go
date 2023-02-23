package handlers

import (
	"bytes"
	"encoding/json"
	"frontend/pkg/render"
	"log"
	"net/http"
	"strings"
)

// Home is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) Leave(w http.ResponseWriter, r *http.Request) {

	// Collect the context from middleware
	myc := r.Context().Value(httpContext).(httpContextStruct)

	if !strings.HasPrefix(r.URL.Path, "/leave/") {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}

	if !myc.Auth {
		var empty any
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
		return
	}

	switch r.Method {
	case "GET":
		leaveGET(w, myc, r.URL.Path)
	// case "POST":
	// 	leavePOST(w, myc, r.URL.Path)
	default:
		log.Println("Sorry, only GET method are supported.")
	}
}

func leavePOST(w http.ResponseWriter, myc httpContextStruct, url string) {
	sURL := strings.Split(url, "/")
	action := sURL[2]

	switch action {
	case "add":
		showLeaveAdd(w, myc)
	case "application":
		showLeaveHome(w, myc)
	default:
		showLeaveHome(w, myc)
	}
}

func leaveGET(w http.ResponseWriter, myc httpContextStruct, url string) {
	sURL := strings.Split(url, "/")
	action := sURL[2]

	switch action {
	case "add":
		showLeaveAdd(w, myc)
	case "application":
		showLeaveHome(w, myc)
	default:
		showLeaveHome(w, myc)
	}
}

func showLeaveHome(w http.ResponseWriter, myc httpContextStruct) {

	var td any

	// email := []byte(myc.User.Email)

	// // getLeave from leave-service
	// resp, err := getLeaveApplications(email)
	// if err != nil {
	// 	return
	// }

	// // Template Data holding the getLeave
	// td := templateData{
	// 	Nickname: myc.User.Nickname,
	// 	Data:     resp,
	// }

	// log.Println("Template Data ", td.Data)

	// Render the page
	if myc.Auth {
		render.RenderTemplate(w, "user.leave.page.gohtml", td)
	} else {
		var empty any
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
	}
}

func showLeaveAdd(w http.ResponseWriter, myc httpContextStruct) {
	switch myc.User.Role {
	case 1:
		var empty any
		render.RenderTemplate(w, "admin.leave.add.page.gohtml", empty)
	case 2:
		var empty any
		render.RenderTemplate(w, "admin.leave.add.page.gohtml", empty)
	default:
		var empty any
		render.RenderTemplate(w, "public.unauthorized.page.gohtml", empty)
	}
}

func getLeaveApplications(encoded []byte) (jsonResponse, error) {
	leaveServiceURL := "http://localhost:8881/getAllLeaveApplications"
	log.Println("I'm going to call...")

	response, err := http.Post(leaveServiceURL, "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		log.Print("Error!!!", err)
		return jsonResponse{}, err
	}
	defer response.Body.Close()

	var payload jsonResponse
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&payload)
	if err != nil {
		log.Print("Error2!!!", err)
		return jsonResponse{}, err
	}
	return payload, nil
}
