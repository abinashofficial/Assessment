package forms

import (
	"Assessment/model"
)

type Repository interface {
	Create(req map[string]string) (model.ConvertedRequest, error)
}
