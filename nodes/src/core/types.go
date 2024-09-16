package core

type GematriaValues struct {
	Kind       int
	ShowName   string
	ValuesName string
}
type GematriaValue struct {
	Rune   string
	Value  int
	Should bool
}
type Gematrias struct {
	Name     string
	ShowName string
	Sum      int
	Values   []GematriaValue
}
