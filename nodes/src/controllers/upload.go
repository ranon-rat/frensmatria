package controllers

import (
	"net/http"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("word-input")
	if input == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	format := core.FormatGematria(core.CalculateAllGematrias(input))
	db.AddGematria(input, format)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
