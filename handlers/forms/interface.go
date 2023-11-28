package forms

import "net/http"

type FormHandlers interface {
	HandlerCreate(w http.ResponseWriter, r *http.Request)
}
