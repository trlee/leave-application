package handlers

import (
	"frontend/pkg/render"
	"net/http"
)

// Home is the handler func executed after middleware validation
// it holds a context sent by the middleware
func (rep *Repository) Testing(w http.ResponseWriter, r *http.Request) {
	myc := r.Context().Value(httpContext).(httpContextStruct)
	myc.authLeave(w)
}

func (r *httpContextStruct) authLeave(w http.ResponseWriter) {
	authorized := r.Auth
	email := r.User.Email
	if authorized {
		render.RenderTemplate(w, "testing.page.gohtml", email)
	} else {
		var empty any
		render.RenderTemplate(w, "unauthorized.page.gohtml", empty)
	}
}
