package core

import "fmt"

func CalculateAllGematrias(input string) []Gematrias {
	out := []Gematrias{}
	for n, g := range GematriasVals {
		values := GematriaVals[g.ValuesName]
		switch g.Kind {
		case sumGematria:
			values, sum := GeneralAdditionGematriaCalculator(input, values)
			out = append(out, Gematrias{
				Name:   n,
				Sum:    sum,
				Values: values,
			})
		case fractalGematria:
			values, sum := GeneralFractalGematriaCalculator(input, values)
			out = append(out, Gematrias{
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
func GeneralFractalGematriaCalculator(input string, gematriaValues map[rune]int) ([]GematriaValue, int) {
	values := make([]GematriaValue, len(input))
	s := 0
	for i, v := range input {
		// i just get the numerical value and if it exists i just use it
		vn, e := gematriaValues[v]
		if !e {
			continue
		}
		s += vn * (i + 1)
		values[i].Value = vn * (i + 1)
		values[i].Rune = string(v)
		values[i].Should = true
	}
	return values, s

}

// this will be used for cool important stuff
func GeneralAdditionGematriaCalculator(input string, gematriaValues map[rune]int) ([]GematriaValue, int) {
	values := make([]GematriaValue, len(input))
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
