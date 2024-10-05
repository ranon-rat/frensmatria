package db

import (
	"database/sql"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func QueryIDGematria(db *sql.DB, id int) (*sql.Rows, error) {
	var query string
	var args []interface{}

	if id != 0 {
		query = `SELECT id, input, search FROM gematrias WHERE id <= ? ORDER BY datePost DESC LIMIT ?`
		args = []interface{}{id, LIMIT}
	} else {
		query = `SELECT id, input, search FROM gematrias ORDER BY datePost DESC LIMIT ?`
		args = []interface{}{LIMIT}
	}

	return db.Query(query, args...)
}

// this is just for the coutn
func Count() (quantity int) {
	db := Connect()
	defer db.Close()
	db.QueryRow(`SELECT COUNT(*) FROM gematrias  ORDER BY datePost DESC`).Scan(&quantity)
	return
}

// so this is just for searching the gematria and other stuff
// in case that the user is interested in a specific kind of gematria it will be setted
func GematriaByID(id int) (tableRows [][]string, lastID int) {

	db := Connect()
	defer db.Close()

	rows, _ := QueryIDGematria(db, id)
	for rows.Next() {
		var inputString, formatGematria string
		// so i just need to scan all of this
		rows.Scan(&lastID, &inputString, &formatGematria)
		tableRows = append(tableRows, append(
			[]string{inputString},
			core.DecodeFGematrias(formatGematria)...,
		))

	}
	return
}
