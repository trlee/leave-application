package handlers

import (
	"authentication/pkg/token"
	"encoding/json"
	"log"
	"net/http"
)

// VerifyCookie verifies the cookie authenticity and validity (expired date)
// if successfull, it returns http.StatusOK and user information
// if failed, it returns http.StatusUnauthorized
func (rep *Repository) VerifyCookie(w http.ResponseWriter, r *http.Request) {
	// Response to front-end
	var verifyCookieResponse = VerifyCookieResponse{}

	// Get Secure Cookie from post request
	// Invoke ParseForm before read form values
	r.ParseForm()
	sc := r.FormValue("cookie")

	// Decode encoded cookie (hold paseto Token)
	t, err := token.DecodeCookie(sc)
	if err != nil {
		log.Fatalf("Request Failed: %s", err)
	}
	// Decode paseto into token.Payload
	p, err := rep.App.Paseto.DecryptToken(t)
	if err != nil {
		log.Fatalf("Request Failed: %s", err)
	}
	// Check if token expired and returns duration
	d, err := p.Valid()
	if err != nil {
		log.Printf("Token expired: %s", err)
		verifyCookieResponse = VerifyCookieResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized connexion, token is expired",
		}
	} else {
		log.Printf("the token is valid for: %s", d)
		verifyCookieResponse = VerifyCookieResponse{
			Status:       http.StatusOK,
			Message:      "authorized connexion",
			TokenPayload: p,
		}
	}

	// TODO autorefresh token & cookie
	// Check when duration is about to end and re-generate it

	json.NewEncoder(w).Encode(&verifyCookieResponse)

}
