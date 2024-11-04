package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

var (
	upgrade = websocket.Upgrader{}
	clients = make(map[*websocket.Conn]bool)
	//Message[Name of channel] message json structure
	Message = make(chan core.Messages)
)

type Pagination struct {
	CurrentPage int
	TotalPages  int
	Back        int
	Next        int
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
