package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type GematriaIndexSearch struct {
	Search               bool
	NotInt               bool
	SimpleGematriaValues []GematriaValue
	SimpleGematriaSum    int
}

func CalculateGematria(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
		r.ParseForm()
		input := strings.TrimSpace(r.Form.Get("word-input"))
		_, err := strconv.Atoi(input)
		NotInt := err != nil
		var simpleV []GematriaValue
		simpleGematriaSum := 0
		if NotInt {
			simpleV, simpleGematriaSum = SimpleGematria(input)
		}
		sent(w, indexT, GematriaIndexSearch{
			Search:               true,
			NotInt:               NotInt,
			SimpleGematriaValues: simpleV,
			SimpleGematriaSum:    simpleGematriaSum,
		})
	default:
		sent(w, indexT, GematriaIndexSearch{})
	}

}
