package db

import (
	"database/sql"
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func QueryIDGematriaSearch(db *sql.DB, sum, kind string, id int) (*sql.Rows, error) {
	like := fmt.Sprintf("%%%s:%s;%%", kind, sum)

	var query string
	var args []interface{}

	if id != 0 {
		query = `SELECT id, input, search FROM gematrias WHERE search LIKE ? AND id <= ? ORDER BY datePost DESC LIMIT ?`
		args = []interface{}{like, id, LIMIT}
	} else {
		query = `SELECT id, input, search FROM gematrias WHERE search LIKE ? ORDER BY datePost DESC LIMIT ?`
		args = []interface{}{like, LIMIT}
	}

	return db.Query(query, args...)
}

// this is just for the coutn
func SearchCount(sum, kind string) (quantity int) {
	db := Connect()
	defer db.Close()
	like := fmt.Sprintf("%%%s:%s;%%", kind, sum)
	db.QueryRow(`SELECT COUNT(*) FROM gematrias WHERE LIKE ?1 ORDER BY datePost DESC`, like).Scan(&quantity)
	return
}

// so this is just for searching the gematria and other stuff
// in case that the user is interested in a specific kind of gematria it will be setted
func SearchGematriaByID(sum, kind string, id int) (tableRows [][]string, lastID int) {
	db := Connect()
	defer db.Close()
	rows, _ := QueryIDGematriaSearch(db, sum, kind, id)

	defer rows.Close()

	for rows.Next() {
		var inputString, formatGematria string
		rows.Scan(&lastID, &inputString, &formatGematria)

		tableRows = append(tableRows, append([]string{inputString}, core.DecodeFGematrias(formatGematria)...))
	}

	return
}
