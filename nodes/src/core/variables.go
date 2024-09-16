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
		"simple": {
			Kind:       sumGematria,
			ValuesName: "simple"},
		"fractal": {
			Kind:       fractalGematria,
			ValuesName: "simple"},
	}
	GematriaVals = map[string]map[rune]int{
		"simple": {
			'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10,
			'k': 20, 'l': 30, 'm': 40, 'n': 50, 'o': 60, 'p': 70, 'q': 80, 'r': 90, 's': 100,
			't': 200, 'u': 300, 'v': 400, 'w': 500, 'x': 600, 'y': 700, 'z': 800,
		},
	}
)
