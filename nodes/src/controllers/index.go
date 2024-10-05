package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

// Template thing
type GematriaIndexSearch struct {
	Search       bool
	NotInt       bool
	SearchInput  string
	GematriaList []core.Gematrias
	OrderTable   []string
}

func Index(w http.ResponseWriter, r *http.Request) {

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
	sent(w, indexT, GematriaIndexSearch{
		Search:       true,
		NotInt:       NotInt,
		SearchInput:  input,
		GematriaList: gematria,
		OrderTable:   core.GematriasOrder,
	})

}
