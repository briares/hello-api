// Package faas for exposing the service as a cloud function.
package faas

import (
	"net/http"

	"github.com/briares/hello-api/handlers/rest"
	"github.com/briares/hello-api/translation"
)

// Translate request handler exposed as a cloud function.
func Translate(w http.ResponseWriter, r *http.Request) {
	translateHandler := rest.NewTranslateHandler(translation.NewStaticService())

	translateHandler.TranslateHandler(w, r)
}
