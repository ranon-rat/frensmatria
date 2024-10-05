package controllers

import (
	"net/http"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func Recent(w http.ResponseWriter, r *http.Request) {
	sent(w, recentT, GeneralGematria{
		GetGematriaSearch: db.GetGematriaPagination(0),
		PagesP:            calculatePagination(0, db.Count()),
		OrderTable:        core.GematriasOrder,
	})
}
