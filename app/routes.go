package app

import (
	"Assessment/handlers/forms"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func runServer(envPort string, h forms.Store) {
	r := mux.NewRouter()
	r.HandleFunc("/public/forms", h.FormHandler.HandlerCreate).Methods(http.MethodPost)

	fmt.Printf("Server listening on port %d...\n", envPort)
	log.Fatal(http.ListenAndServe(":"+envPort, r))
}
