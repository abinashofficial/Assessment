package forms

import (
	"Assessment/consts"
	"Assessment/model"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

func New(db model.MongoClient) Repository {
	repo := db.PrimaryClient.Database(consts.SupportPortalDB).Collection(consts.FormsColl)
	secondaryRepo := db.SecondaryClient.Database(consts.SupportPortalDB).Collection(consts.FormsColl)
	return &formsRepo{
		repo,
		secondaryRepo,
	}
}

type formsRepo struct {
	repo          *mongo.Collection
	secondaryRepo *mongo.Collection
}

func (s *formsRepo) Create(req map[string]string) (model.ConvertedRequest, error) {
	// Use a wait group to wait for goroutines to finish
	var wg sync.WaitGroup
	var forms = model.ConvertedRequest{}
	var err error

	if req == nil {
		return model.ConvertedRequest{}, fmt.Errorf("empty map")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Send the map through the channel
		model.RequestChannel <- req
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		// Receive the map from the channel
		forms, err = worker(model.RequestChannel)
		if err != nil {
			fmt.Println("Error converting numeric part to integer:", err)
		}
	}()

	wg.Wait()

	//close(model.RequestChannel)

	return forms, err
}
