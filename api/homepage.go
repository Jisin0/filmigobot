package api

import (
	"net/http"
)

// Serve index.html or homepage.
func ServeHomepage(w http.ResponseWriter, r *http.Request) {
	// Serve the static file using ServeFile
	http.ServeFile(w, r, "index.html")
}
