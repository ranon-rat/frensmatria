package controllers

import (
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

const (
	sumGematria = 0
)

var (
	GematriasVals = map[string]core.GematriaValues{
		"simple": {
			Kind: sumGematria,
			Values: map[rune]int{
				'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10,
				'k': 20, 'l': 30, 'm': 40, 'n': 50, 'o': 60, 'p': 70, 'q': 80, 'r': 90, 's': 100,
				't': 200, 'u': 300, 'v': 400, 'w': 500, 'x': 600, 'y': 700, 'z': 800,
			}},
	}
)

func CalculateAllGematrias(input string) []core.Gematrias {
	out := []core.Gematrias{}
	for n, g := range GematriasVals {
		switch g.Kind {
		case sumGematria:
			values, sum := GeneralAdditionGematriaCalculator(input, g.Values)
			out = append(out, core.Gematrias{
				Name:   n,
				Sum:    sum,
				Values: values,
			})
		default:
			fmt.Println("weird")

		}

	}

	return out
}

// this will be used for cool important stuff
func GeneralAdditionGematriaCalculator(input string, gematriaValues map[rune]int) ([]core.GematriaValue, int) {
	values := make([]core.GematriaValue, len(input))
	s := 0
	for i, v := range input {
		// i just get the numerical value and if it exists i just use it
		vn, e := gematriaValues[v]
		if !e {
			continue
		}
		values[i].Value = vn
		s += gematriaValues[v]
		values[i].Rune = string(v)
		values[i].Should = true
	}
	return values, s
}
