package service

import "Assessment/model"

type ServiceRepository interface {
	Create(service *model.Service) error
	Update(service *model.Service) error
	FindByID(id string) (*model.Service, error)
	Toggle(service *model.Service) error
}
