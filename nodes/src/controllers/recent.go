package controllers

import (
	"net/http"
	"strconv"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func Recent(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	sent(w, recentT, GeneralGematria{
		Render:            false,
		GetGematriaSearch: db.GetGematriaPagination(page - 1),
		PagesP:            calculatePagination(page, db.Count()),
		OrderTable:        core.GematriasOrder,
	})
}
