package handlers

import "authentication/pkg/token"

// RPCPayload is a struct which holds the information to be sent to the logger-service
type RPCPayload struct {
	Collection string
	Name       string
	Data       string
}

// AuthPayload is the embedded type (in RequestPayload) that describes an authentication request
type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// jsonResponse is a struct which holds the response to be sent back
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// VerifyCookieResponse is the response send back to front-end
// that describes the cookie & token validity
type VerifyCookieResponse struct {
	Status       int              `json:"status"`
	Message      string           `json:"message"`
	TokenPayload *token.TokenData `json:"tokenpayload"`
}

type LeavePayload struct {
	Email    string `json:"email"`
	Duration int    `json:"duration"`
	Reason   string `json:"reason"`
}
