package handlers

import (
	"leave/pkg/config"
)

var (
	httpContext httpContextStruct
	Repo        *Repository
)

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
