package core

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func LogColor(input ...any) {
	c := color.New(color.Bold).AddRGB(0, 255, 0).SprintFunc()
	t := time.Now()

	fmt.Println(append([]any{
		c(fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d ",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second()))}, input...)...)

}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GematriaSharing2Base64(g GematriaSharing) string {
	b, _ := json.Marshal(g)
	return base64.RawStdEncoding.EncodeToString(b)
}
func Base64_2GematriaSharing(b string) GematriaSharing {
	var g GematriaSharing
	d, _ := base64.RawStdEncoding.DecodeString(b)
	json.Unmarshal(d, &g)
	return g
}

func CalculateAllGematrias(input string) []Gematrias {
	out := []Gematrias{}
	// so this will follow a range
	for _, n := range GematriasOrder {

		g := GematriasVals[n]
		i := input
		if !g.Upper {
			i = strings.ToLower(i)
		}
		values := GematriaVals[g.ValuesName]
		switch g.Kind {
		case sumGematria:
			values, sum := GeneralAdditionGematriaCalculator(i+" ", values)
			out = append(out, Gematrias{
				Name:     n,
				Sum:      sum,
				ShowName: g.ShowName,
				Values:   values,
			})
		case fractalGematria:
			values, sum := GeneralFractalGematriaCalculator(i+" ", values)
			out = append(out, Gematrias{
				Name:     n,
				ShowName: g.ShowName,
				Sum:      sum,
				Values:   values,
			})
		default:
			fmt.Println("weird")

		}

	}

	return out
}

// now lets add more stuff here
// this will encode and make a format for the gematrias
// this is for storing the values and other stuff, its important to keep it
func FormatGematria(gematrias []Gematrias) (out string) {

	for _, v := range gematrias {
		out += fmt.Sprintf("%s:%d;", v.Name, v.Sum)
	}
	return
}

// this will decode the formated gematrias
func DecodeFGematrias(encoded string) (out []string) {
	formatedGematrias := strings.Split(encoded, ";")
	// why i am following the list of gematriaorder when this is not really required?
	// well i dont want to scan 2 variables kek
	for i, v := range GematriasOrder {

		sum := ""
		fmt.Sscanf(formatedGematrias[i], v+":%s", &sum)
		out = append(out, sum)

	}
	return
}

// this is just generalized gematria calculations

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
