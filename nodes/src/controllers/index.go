package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	input := strings.TrimSpace(r.URL.Query().Get("word-input"))
	if input == "" {
		sent(w, indexT, GematriaIndexSearch{})
		return
	}
	_, err := strconv.Atoi(input)
	NotInt := err != nil
	gematria := []core.Gematrias{}
	if NotInt {
		gematria = core.CalculateAllGematrias(input)
	}
	// this is just for the table

	results := [][]string{}
	itemsCount := 0

	if !NotInt {
		results = db.SearchGematriaPaginated(input, core.GematriasOrder[0], page-1)
		itemsCount = db.SearchCount(input, core.GematriasOrder[0])
	} else {
		results = db.SearchGematriaPaginated(strconv.Itoa(gematria[0].Sum), gematria[0].Name, page-1)
		itemsCount = db.SearchCount(strconv.Itoa(gematria[0].Sum), gematria[0].Name)

	}
	fmt.Println(calculatePagination(page, itemsCount), itemsCount, db.SearchCount(input, core.GematriasOrder[0]))
	sent(w, indexT, GematriaIndexSearch{
		Search:            true,
		NotInt:            NotInt,
		SearchInput:       input,
		GematriaList:      gematria,
		OrderTable:        core.GematriasOrder,
		GetGematriaSearch: results,

		PagesP: calculatePagination(page, itemsCount),
	})

}
