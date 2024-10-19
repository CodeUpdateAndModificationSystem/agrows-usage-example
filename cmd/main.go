package main

import (
	"agrows-usage-example/internal/middleware"
	"agrows-usage-example/internal/socket"
	"agrows-usage-example/web"
	"net/http"

	log "github.com/dikkadev/dnutlogger"

	"github.com/a-h/templ"
)

func main() {
	log.SetMinLevel(log.DEBUG)

	router := http.NewServeMux()

	var fileServer http.Handler
	fileServer = http.FileServer(http.Dir("web/assets/dist"))
	router.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/" {
			templ.Handler(web.Index()).ServeHTTP(w, r)
		} else {
			http.Error(w, "404 not found", http.StatusNotFound)
		}
	})

	router.HandleFunc("/api/ws", socket.WebSocketHandler)

	log.Info("Starting server in dev mode on port 8080\n")
	err := http.ListenAndServe("127.0.0.1:8080", middleware.Logging(router))
	if err != nil {
		log.Errorf(true, "Error starting server: %v\n", err)
	}
}
