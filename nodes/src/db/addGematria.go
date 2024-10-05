package db

import (
	"fmt"
	"time"
)

// so this will store the gematria in the databsae, its not that complicated obviously xd
func AddGematria(input, format string) {
	query := `INSERT INTO gematrias(input,search,datePost) VALUES(?1,?2,?3)`
	db := Connect()
	defer db.Close()
	_, err := db.Exec(query, input, format, int(time.Now().Unix()))
	if err != nil {
		fmt.Println(err)
	}
	// no deberia de ser muy complicado
	// ahora voy a intentar correr una cosa para hacer el setup
}
