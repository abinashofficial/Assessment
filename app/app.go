package app

import (
	"Assessment/handlers/forms"
	"Assessment/log"
	"Assessment/model"
	forms2 "Assessment/services/forms"
	mongo_store "Assessment/store"
	forms_repo "Assessment/store/forms"
	"Assessment/store/mongo"
	"Assessment/tapcontext"
	"context"
	"os"
)

var h forms.Store
var service forms2.Store
var repos mongo_store.Store

func setupHandlers() {
	h = forms.Store{
		FormHandler: forms.New(service),
	}
}

func setupServices() {
	formServ := forms2.New(repos)

	service = forms2.Store{
		FormServ: formServ,
	}

}

func setupRepos(client model.MongoClient) {

	repos = mongo_store.Store{
		FormsRepo: forms_repo.New(client),
	}

}

func Start() {
	ctx := tapcontext.TContext{
		Context:    context.Background(),
		TapContext: tapcontext.TapContext{},
	}

	envPort := os.Getenv("PORT")
	mongoURL := os.Getenv("MONGO_URL")
	secondaryMongoURL := os.Getenv("SECONDARY_MONGO_URL")

	if envPort == "" {
		envPort = "4997"
	}
	//allMongoURL := model.MongoURL{mongoURL, secondaryMongoURL}
	allMongoURL := map[string]string{
		"mongoURL":          mongoURL,
		"secondaryMongoURL": secondaryMongoURL,
	}
	client, mongoErr := mongo.Init(ctx, allMongoURL)
	defer mongo.Disconnect(ctx)

	if mongoErr.PrimaryClientError != nil {
		log.GenericError(ctx, mongoErr.PrimaryClientError, log.FieldsMap{"error": "MongoURL Connection Failed"})
		log.FatalLog(ctx, mongoErr.PrimaryClientError, nil)
	}

	setupServices()
	setupHandlers()
	setupRepos(client)
	runServer(envPort, h)

}
