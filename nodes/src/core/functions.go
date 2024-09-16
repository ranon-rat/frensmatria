package core

import "fmt"

func CalculateAllGematrias(input string) []Gematrias {
	out := []Gematrias{}
	for n, g := range GematriasVals {
		values := GematriaVals[g.ValuesName]
		switch g.Kind {
		case sumGematria:
			values, sum := GeneralAdditionGematriaCalculator(input+" ", values)
			out = append(out, Gematrias{
				Name:   n,
				Sum:    sum,
				Values: values,
			})
		case fractalGematria:
			values, sum := GeneralFractalGematriaCalculator(input+" ", values)
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
	values := []GematriaValue{}
	s := 0
	spaceS := 0
	lastTrue := true
	trueValue := 0
	for _, v := range input {
		var r GematriaValue
		// i just get the numerical value and if it exists i just use it
		vn, e := gematriaValues[v]
		if !e {
			if lastTrue {
				continue
			}
			r.Value = spaceS
			r.Should = false
			values = append(values, r)
			spaceS = 0
			lastTrue = true
			continue
		}
		k := vn * (trueValue + 1)
		s += k
		spaceS += k
		r.Value = k
		r.Rune = string(v)
		r.Should = true
		values = append(values, r)
		lastTrue = false
		trueValue++
	}
	return values, s

}

// this will be used for cool important stuff
func GeneralAdditionGematriaCalculator(input string, gematriaValues map[rune]int) ([]GematriaValue, int) {
	values := []GematriaValue{}
	s := 0
	spaceS := 0
	lastTrue := true
	for _, v := range input {
		var r GematriaValue
		// i just get the numerical value and if it exists i just use it
		vn, e := gematriaValues[v]
		if !e {
			if lastTrue {
				continue
			}
			r.Value = spaceS
			r.Should = false
			values = append(values, r)
			spaceS = 0
			lastTrue = true
			continue
		}
		s += vn
		spaceS += vn
		r.Value = vn
		r.Rune = string(v)
		r.Should = true
		values = append(values, r)
		lastTrue = false

	}
	return values, s
}
