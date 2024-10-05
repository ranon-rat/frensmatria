package db

import (
	"time"
)

// so this will store the gematria in the databsae, its not that complicated obviously xd
// if there is an error obviously its because the value its repeated
func AddGematria(input, format string) (err error) {
	query := `INSERT INTO gematrias(input,search,datePost) VALUES(?1,?2,?3)`
	db := Connect()
	defer db.Close()
	_, err = db.Exec(query, input, format, int(time.Now().Unix()))
	if err != nil {
		return
	}
	return
}
