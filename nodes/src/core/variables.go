package core

// pagination shit
const LIMIT = 15

// gematria shit
const (
	sumGematria     = iota
	fractalGematria = iota
)

// gematria kind of shit
var (

	/// this is for the table and the output, i want some consistency so this is quite important
	GematriasOrder = []string{
		"AQ",
		"synx",
	}
	// this is just generalized
	GematriasVals = map[string]GematriaValues{

		"AQ": {
			Kind:       sumGematria,
			ShowName:   "Alpha numeric Qabbalah",
			ValuesName: "AQ",
			Upper:      false,
		},
		"synx": {
			Kind:       sumGematria,
			ShowName:   "Synx",
			ValuesName: "synx",

			Upper: false,
		},
	}
	// they dont really need to have some order
	GematriaVals = map[string]map[rune]int{
		"synx": {
			'0': 1, '1': 2, '2': 3, '3': 4, '4': 5, '5': 6, '6': 7, '7': 9, '8': 10, '9': 12,
			'a': 14, 'b': 15, 'c': 18, 'd': 20, 'e': 21, 'f': 28, 'g': 30, 'h': 35,
			'i': 36, 'j': 42, 'k': 45, 'l': 60, 'm': 63, 'n': 70, 'o': 84, 'p': 90, 'q': 105, 'r': 126,
			's': 140, 't': 180, 'u': 210, 'v': 252, 'w': 315, 'x': 420, 'y': 630, 'z': 1260,
		},
		"AQ": {
			'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
			'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15, 'g': 16, 'h': 17,
			'i': 18, 'j': 19, 'k': 20, 'l': 21, 'm': 22, 'n': 23, 'o': 24, 'p': 25, 'q': 26, 'r': 27,
			's': 28, 't': 29, 'u': 30, 'v': 31, 'w': 32, 'x': 33, 'y': 34, 'z': 35,
		},
	}
)
