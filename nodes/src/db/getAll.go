package db

import (
	"fmt"
	"sync"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func Count() (quantity int) {
	db := Connect()
	defer db.Close()
	db.QueryRow(`SELECT COUNT(*) FROM gematrias `).Scan(&quantity)
	return
}
func GetLastDate() (date int) {
	query := `SELECT  datePost FROM gematrias TOP LIMIT 1;`
	db := Connect()
	defer db.Close()
	db.QueryRow(query).Scan(&date)
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

// Only to be used in the connections folder
// i am thinking that i could actually just use channels for communicating with other stuff
// but, why i would do that?
// this is simple, it doesnt needs more
func GetAllGematria(conn *webrtc.DataChannel, date int) {
	// for some reason the > operator its working as a >= its quite weird
	query := `SELECT input, datePost From gematrias WHERE datePost>?1 AND datePost<>?1`
	db := Connect()
	defer db.Close()
	rows, _ := db.Query(query, date)
	// we are back
	var wg sync.WaitGroup

	for rows.Next() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g := core.GematriaSharing{}
			rows.Scan(&g.Content, &g.Date)
			conn.SendText(fmt.Sprintf("compare %s", core.GematriaSharing2Base64(g)))
		}()
	}
	wg.Wait()

	// i hope i dont have any problem because of the concurrency XD
	conn.SendText("end")
}
