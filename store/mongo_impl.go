package store

import (
	"Assessment/consts"
	"Assessment/model"
	"Assessment/tapcontext"
	"go.mongodb.org/mongo-driver/mongo"
)

type formsRepo struct {
	repo          *mongo.Collection
	secondaryRepo *mongo.Collection
}

func New(db model.MongoClient) Repository {
	repo := db.PrimaryClient.Database(consts.SupportPortalDB).Collection(consts.FormsColl)
	secondaryRepo := db.SecondaryClient.Database(consts.SupportPortalDB).Collection(consts.FormsColl)
	return &formsRepo{
		repo:          repo,
		secondaryRepo: secondaryRepo,
		//cacheRepo:     cacheRepo,
		//redisRequired: redisRequired,
	}
}

func (r *formsRepo) GetAllForms(ctx tapcontext.TContext) error {
	panic("dhj")
}
