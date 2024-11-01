package router

import (
	"log"
	"net/http"

	"github.com/ranon-rat/frensmatria/nodes/src/controllers"
)

func Setup(port string) {
	log.Println("starting server")
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs)) // basic files
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/upload", controllers.Upload)
	http.HandleFunc("/recent", controllers.Recent)

	// so this is a setup
	controllers.SetupTemplate()
	// nothing to see here just protocolarial stuff

	http.ListenAndServe(":"+port, nil)
}
