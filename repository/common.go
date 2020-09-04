package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB


func InitiateDB(dbInstance *sql.DB) {
	db = dbInstance
}
