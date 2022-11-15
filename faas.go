// Package faas for exposing the service as a cloud function.
package faas

import (
	"net/http"

	"github.com/briares/hello-api/handlers/rest"
)

// Translate request handler exposed as a cloud function.
func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
