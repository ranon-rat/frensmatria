package core

type GematriaValues struct {
	Kind       int
	ValuesName string
}
type GematriaValue struct {
	Rune   string
	Value  int
	Should bool
}
type Gematrias struct {
	Name   string
	Sum    int
	Values []GematriaValue
}
