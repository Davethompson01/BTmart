package config

import (
	"database/sql"
	"log/slog"
)

type Apiconfig struct {
	DB     *sql.DB
	Logger *slog.Logger
}
