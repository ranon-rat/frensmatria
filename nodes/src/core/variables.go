package core

// connection shit
const (
	Nothing   = 0
	ConnectTo = 1
	Confirm   = 2
)

// gematria shit
const (
	sumGematria = 0
)

var (
	GematriasVals = map[string]GematriaValues{
		"simple": {
			Kind: sumGematria,
			Values: map[rune]int{
				'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10,
				'k': 20, 'l': 30, 'm': 40, 'n': 50, 'o': 60, 'p': 70, 'q': 80, 'r': 90, 's': 100,
				't': 200, 'u': 300, 'v': 400, 'w': 500, 'x': 600, 'y': 700, 'z': 800,
			}},
	}
)
