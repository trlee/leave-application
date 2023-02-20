package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// ResponseCookieVerify {
// 	Status       int       `json:"status"`
// 	Message      string    `json:"message"`
// 	TokenPayload TokenData `json:"tokenpayload"`
// }

// type TokenData struct {
// 	ID        uuid.UUID `json:"uuid"`
// 	Email     string    `json:"email"`
// 	Nickname  string    `json:"nickname"`
// 	Role      int       `json:"role"`
// 	IssuedAt  time.Time `json:"issue_at"`
// 	ExpiredAt time.Time `json:"expired_at"`
// }

// Middleware checks if secure cookie exists
// if the cookie is valid, send it to authentication-service (resp = ResponseCookieVerify)
// if resp.Status = http.StatusAccepted (200) set httpContext to request and forward it
// if resp.Status !=  http.StatusAccepted redirect to unauthorized page
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get secure cookie from httponly header request
		cookie, err := r.Cookie(CookieName)
		if err == nil {
			// Send encoded cookie to authentication-service for validation
			verifyCookiePayload := VerifyCookiePayload{
				Cookie: cookie.Value,
			}
			// Get authentication-service response
			// returns type ResponseCookieVerify
			resp, err := verifyCookiePayload.verifyCookie()
			if err != nil {
				log.Fatalf("Request Failed: %s", err)
			}

			// Authorized request
			if resp.Status == 200 {
				// httpContext = httpContextStruct{}
				// Set httpContextStruct
				r = r.WithContext(context.WithValue(
					r.Context(),
					httpContext,
					httpContextStruct{
						Auth: true,
						User: resp.TokenPayload,
					},
				))
				next.ServeHTTP(w, r)
			} else {
				// Unauthorized request
				http.Redirect(w, r, "http://localhost/unauthorized", http.StatusSeeOther)
			}

		} else {
			// Unauthorized request
			http.Redirect(w, r, "http://localhost/unauthorized", http.StatusSeeOther)
		}

	})
}

// Verify cookie's authorization and validity
// send post request to authentication-service
func (p *VerifyCookiePayload) verifyCookie() (ResponseCookieVerify, error) {
	params := url.Values{}
	params.Add("cookie", p.Cookie)

	resp, err := http.PostForm("http://localhost:8081/verifyCookie", params)
	if err != nil {
		log.Fatalf("Request Failed: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Log the request body
	if err != nil {
		log.Fatalf("Request Failed: %s", err)
	}

	respAPI := ResponseCookieVerify{}
	err = json.Unmarshal(body, &respAPI)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	return respAPI, nil
}
