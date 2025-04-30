package service
import(
	"Assessment/model"
)

type ServiceService interface {
	CreateService(service *model.Service) error
	UpdateService(id string, data *model.Service) (*model.Service, error)
	ToggleService(id string) (*model.Service, error)
}