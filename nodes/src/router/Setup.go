package router

import (
	"net/http"
	"os"

	"github.com/ranon-rat/frensmatria/nodes/src/controllers"
)

func Setup() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs)) // basic files
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/upload", controllers.Upload)
	http.HandleFunc("/recent", controllers.Recent)

	// so this is a setup
	controllers.SetupTemplate()
	// nothing to see here just protocolarial stuff
	port, e := os.LookupEnv("PORT")
	if !e {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
