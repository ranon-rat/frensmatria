package db

import (
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

// this is just for the coutn
func Count() (quantity int) {
	db := Connect()
	defer db.Close()
	db.QueryRow(`SELECT COUNT(*) FROM gematrias `).Scan(&quantity)
	return
}

// so this is just for searching the gematria and other stuff
// in case that the user is interested in a specific kind of gematria it will be setted
func GetGematriaPagination(page int) (tableRows [][]string) {
	query := `SELECT  input, search FROM gematrias ORDER BY datePost DESC LIMIT ? OFFSET ?`
	db := Connect()
	defer db.Close()

	rows, _ := db.Query(query, core.LIMIT, page*core.LIMIT)
	for rows.Next() {
		var inputString, formatGematria string
		// so i just need to scan all of this
		rows.Scan(&inputString, &formatGematria)
		tableRows = append(tableRows, append(
			[]string{inputString},
			core.DecodeFGematrias(formatGematria)...,
		))
	}

	return
}
