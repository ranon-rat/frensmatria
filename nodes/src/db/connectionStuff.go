package db

import (
	"fmt"
	"sync"

	"github.com/pion/webrtc/v3"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

// this is for the protocol
func GetAllGematria(conn *webrtc.DataChannel, date int) {
	query := `SELECT input, datePost From gematrias WHERE datePost>?`
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
	// esto deberia de ser suficiente creo yo, espero que no genere problemas el enviarlo de manera asincrona
	conn.SendText("end")
}

func GetLastDate() (date int) {
	query := `SELECT  datePost FROM gematrias TOP LIMIT 1;`
	db := Connect()
	defer db.Close()
	db.QueryRow(query).Scan(&date)
	return
}
