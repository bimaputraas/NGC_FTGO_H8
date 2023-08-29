package handler

import "database/sql"

type Handler struct {
	HandlerDB *sql.DB
}