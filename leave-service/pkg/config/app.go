package config

import (
	"database/sql"
	"leave/pkg/pg"
)

type AppConfig struct {
	DB     *sql.DB
	Models pg.Models
}
