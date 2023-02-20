package handlers

import (
	"time"

	"github.com/google/uuid"
)

type RPCPayload struct {
	Collection string
	Name       string
	Data       string
}

type ReportPayload struct {
	Email string `json:"email"`
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type TokenData struct {
	ID        uuid.UUID `json:"uuid"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Role      int       `json:"role"`
	IssuedAt  time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type httpContextStruct struct {
	Auth bool
	User TokenData
}
