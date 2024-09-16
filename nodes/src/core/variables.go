package core

// connection shit
const (
	Nothing   = 0
	ConnectTo = 1
	Confirm   = 2
)

// gematria shit
const (
	sumGematria     = 0
	fractalGematria = 1
)

var (

	// con esto podemos dejar cuales son las que existen, so thats good enough for me and for everyone else
	// rn i have to think on how the fuck iam building the db
	// i could be using a csv, a json format or idk
	// maybe using a simple uhh text divided by commas?
	GematriasVals = map[string]GematriaValues{

		"AQ": {
			Kind:       sumGematria,
			ShowName:   "Alpha numeric Qabbalah",
			ValuesName: "AQ",
		},
		"synx": {
			Kind:       sumGematria,
			ShowName:   "Synx",
			ValuesName: "synx"},
	}
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
