package handlers

import (
	"frontend/pkg/render"
	"net/http"
)

// Home is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) Leave(w http.ResponseWriter, r *http.Request) {
	// myc := r.Context().Value(httpContext).(httpContextStruct)
	// all, err := rep.App.Models.LeaveApplication.GetAllApplications(myc.User.Email)
	// all2, err := rep.App.Models.LeaveApplication.GetAllApplications(myc.User.Email)
	// if err != nil {
	// 	log.Println("error while fetching all applications", err)
	// }

	// var td TemplateData
	// td.User = myc.User
	// td.Data = all

	// render.RenderTemplate(w, "base-leave.application.page.gohtml", td)

	myc := r.Context().Value(httpContext).(httpContextStruct)
	myc.applyLeave(w)
}

func (r *httpContextStruct) applyLeave(w http.ResponseWriter) {
	authorized := r.Auth
	email := r.User.Email
	if authorized {
		render.RenderTemplate(w, "base.leave.apply.page.gohtml", email)
	} else {
		var empty any
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
	}
}
