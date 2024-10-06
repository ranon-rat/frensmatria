package controllers

import (
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
	kind := r.URL.Query().Get("kind")
	if _, e := core.GematriasVals[kind]; !e {
		kind = core.GematriasOrder[0]
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

	var results [][]string
	itemsCount := 0

	if !NotInt {
		results = db.SearchGematriaPaginated(input, kind, page-1)
		itemsCount = db.SearchCount(input, kind)
	} else {
		k := core.Gematrias{}
		for i := 0; i < len(gematria); i++ {
			if gematria[i].Name == kind {
				k = gematria[i]
				break
			}

		}
		results = db.SearchGematriaPaginated(strconv.Itoa(k.Sum), kind, page-1)
		itemsCount = db.SearchCount(strconv.Itoa(k.Sum), kind)

	}
	sent(w, indexT, GematriaIndexSearch{
		Render:            true,
		Search:            true,
		NotInt:            NotInt,
		SearchInput:       input,
		GematriaList:      gematria,
		OrderTable:        core.GematriasOrder,
		GetGematriaSearch: results,
		PagesP:            calculatePagination(page, itemsCount),
	})

}
