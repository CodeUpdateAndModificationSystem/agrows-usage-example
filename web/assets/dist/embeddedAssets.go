package assets

import (
	"embed"
	"net/http"
)

//go:embed *.css *.js *.png
var Assets embed.FS

func ServeAssets(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.FS(Assets))
	http.StripPrefix("/assets/", fs).ServeHTTP(w, r)
}
