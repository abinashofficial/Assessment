package forms

import (
	"Assessment/model"
	"Assessment/store"
	"Assessment/store/forms"
	"fmt"
)

func New(store store.Store) Service {
	return &formServ{
		formsRepo: store.FormsRepo,
	}
}

type formServ struct {
	formsRepo forms.Repository
}

func (s *formServ) Create(req map[string]string) (model.ConvertedRequest, error) {
	form, err := s.formsRepo.Create(req)
	if err != nil {
		fmt.Println("Create Form:", err)
		return model.ConvertedRequest{}, err
	}

	return form, nil
}
