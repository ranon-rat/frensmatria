package core

type GematriaValues struct {
	Kind       int
	ShowName   string
	ValuesName string
	Upper      bool
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

// this will be used for sharing information between nodes :D
type GematriaSharing struct {
	Content string `json:"content"`
	Date    int    `json:"date"`
}

type Messages struct {
	Content   string `json:"content"`
	Timestamp int    `json:"timestamp"`
	Author    string `json:"author"`
}
