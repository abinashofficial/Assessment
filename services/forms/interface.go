package forms

import "Assessment/model"

type Service interface {
	Create(original map[string]string) (model.ConvertedRequest, error)
}
