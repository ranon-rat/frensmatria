package db

import "github.com/ranon-rat/frensmatria/nodes/src/core"

// so this will store the gematria in the databsae, its not that complicated obviously xd
// if there is an error obviously its because the value its repeated
func AddGematria(input string, date int) (err error) {
	query := `INSERT INTO gematrias(input,search,datePost) VALUES(?1,?2,?3)`
	db := Connect()
	defer db.Close()
	_, err = db.Exec(query, input, core.FormatGematria(core.CalculateAllGematrias(input)), date)
	if err != nil {
		return
	}
	return
}
