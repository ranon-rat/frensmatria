package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

// Template thing
type GematriaIndexSearch struct {
	Search       bool
	NotInt       bool
	GematriaList []core.Gematrias
}

func CalculateGematria(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
		r.ParseForm()
		input := strings.TrimSpace(r.Form.Get("word-input"))
		_, err := strconv.Atoi(input)
		NotInt := err != nil
		gematria := []core.Gematrias{}
		if NotInt {
			gematria = core.CalculateAllGematrias(input)
		}
		sent(w, indexT, GematriaIndexSearch{
			Search:       true,
			NotInt:       NotInt,
			GematriaList: gematria,
		})
	default:
		sent(w, indexT, GematriaIndexSearch{})
	}

}
