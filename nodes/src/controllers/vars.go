package controllers

import "github.com/ranon-rat/frensmatria/nodes/src/core"

type Pagination struct {
	CurrentPage int
	TotalPages  int
	Render      bool
	Pages       []PagesElement
}
type PagesElement struct {
	Page    int
	Current bool
}
type GematriaIndexSearch struct {
	Search            bool
	NotInt            bool
	SearchInput       string
	GematriaList      []core.Gematrias
	GetGematriaSearch [][]string
	OrderTable        []string
	PagesP            Pagination
	Render            bool
}
type GeneralGematria struct {
	Render            bool
	GetGematriaSearch [][]string
	OrderTable        []string
	PagesP            Pagination
}
