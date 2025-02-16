package handlers

import (
	"net/http"

	"github.com/codinomello/webjetos-go/views/templates"
)

func HandleTemplIndex(w http.ResponseWriter, r *http.Request) error {
	return templates.Index().Render(r.Context(), w)
}
