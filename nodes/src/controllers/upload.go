package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
	"github.com/ranon-rat/frensmatria/nodes/src/core/channels"
	"github.com/ranon-rat/frensmatria/nodes/src/db"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("word-input")
	if input == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	g := core.GematriaSharing{
		Content: input,
		Date:    int(time.Now().Unix()),
	}
	if db.AddGematria(g.Content, g.Date) == nil {
		channels.SendMessage(fmt.Sprintf("new %s", core.GematriaSharing2Base64(g)), "")
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
