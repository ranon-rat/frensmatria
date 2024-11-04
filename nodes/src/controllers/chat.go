package controllers

import "net/http"

func Chat(w http.ResponseWriter, r *http.Request) {

	sent(w, chatT, nil)
}
