package service

import (
	"Assessment/model"
	"gorm.io/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) Create(service *model.Service) error {
	return r.db.Create(service).Error
}

func (r *serviceRepository) Update(service *model.Service) error {
	return r.db.Save(service).Error
}

func (r *serviceRepository) FindByID(id string) (*model.Service, error) {
	var service model.Service
	if err := r.db.First(&service, id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *serviceRepository) Toggle(service *model.Service) error {
	service.Active = !service.Active
	return r.db.Save(service).Error
}
