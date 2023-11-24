package app

import (
	"Assessment/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func runServer(envPort string, h handlers.Store) {
	r := mux.NewRouter()
	r.HandleFunc("/public/forms", h.FormHandler.HandlerCreate).Methods(http.MethodPost)

	fmt.Printf("Server listening on port %d...\n", envPort)
	log.Fatal(http.ListenAndServe(":"+envPort, r))
}
