package PgSql

import (
	"database/sql"
)

func RunQuery(db *sql.DB, cmd string) (*sql.Rows, error) {
	return db.Query(cmd)
}

func ExecuteCommand(cmd string, db *sql.DB) (*sql.Result, error) {
	r, err := db.Exec(cmd)
	return &r, err
}
