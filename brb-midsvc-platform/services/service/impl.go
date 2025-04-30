package service

import (
	"Assessment/model"
	"Assessment/store/service"
)

type serviceService struct {
	repo service.ServiceRepository
}

func NewServiceService(r service.ServiceRepository) ServiceService {
	return &serviceService{repo: r}
}

func (s *serviceService) CreateService(service *model.Service) error {
	return s.repo.Create(service)
}

func (s *serviceService) UpdateService(id string, data *model.Service) (*model.Service, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	data.ID = existing.ID
	if err := s.repo.Update(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *serviceService) ToggleService(id string) (*model.Service, error) {
	service, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Toggle(service); err != nil {
		return nil, err
	}
	return service, nil
}
