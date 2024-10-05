package db

import (
	"database/sql"
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func QueryGematriaSearch(db *sql.DB, sum, kind string, offset int) (*sql.Rows, error) {
	like := fmt.Sprintf("%%%s:%s;%%", kind, sum)
	query := `SELECT  input, search FROM gematrias WHERE search LIKE ?  ORDER BY datePost DESC LIMIT ? OFFSET ?`
	return db.Query(query, like, core.LIMIT, offset)
}

// this is just for the coutn
func SearchCount(sum, kind string) (quantity int) {
	db := Connect()
	defer db.Close()
	like := fmt.Sprintf("%%%s:%s;%%", kind, sum)
	fmt.Println(like)
	db.QueryRow(`SELECT COUNT(*) FROM gematrias WHERE search LIKE ?1`, like).Scan(&quantity)
	return
}

// so this is just for searching the gematria and other stuff
// in case that the user is interested in a specific kind of gematria it will be setted
func SearchGematriaPaginated(sum, kind string, page int) (tableRows [][]string) {
	db := Connect()
	defer db.Close()
	rows, _ := QueryGematriaSearch(db, sum, kind, page*core.LIMIT)
	defer rows.Close()
	for rows.Next() {
		var inputString, formatGematria string
		rows.Scan(&inputString, &formatGematria)
		tableRows = append(tableRows, append([]string{inputString}, core.DecodeFGematrias(formatGematria)...))
	}

	return
}
