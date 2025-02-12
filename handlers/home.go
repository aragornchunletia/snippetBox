package handlers

import (
	"net/http"

	"snippetbox.aragorn.net/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
