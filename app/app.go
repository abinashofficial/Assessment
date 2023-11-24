package app

import (
	"Assessment/handlers"
	"Assessment/services"
	"Assessment/store"
	"os"
)

var h handlers.Store
var service services.Store
var repos store.Store

func setupHandlers() {
	h = handlers.Store{
		FormHandler: handlers.New(service),
	}
}

func setupServices() {
	formServ := services.New(repos)

	service = services.Store{
		FormServ: formServ,
	}

}

//func setupRepos(client model.MongoClient) {
//
//	repos = store.Store{
//		FormsRepo: store.New(client),
//	}
//
//}

func Start() {
	envPort := os.Getenv("PORT")

	if envPort == "" {
		envPort = "4996"
	}

	setupServices()
	setupHandlers()
	runServer(envPort, h)

}
