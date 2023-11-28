package mock_repos

import (
	"Assessment/model"
	"github.com/stretchr/testify/mock"
)

type MockFormRepo struct {
	mock.Mock
}

func (m MockFormRepo) Create(req map[string]string) (model.ConvertedRequest, error) {
	//TODO implement me
	args := m.Called(req)
	return args.Get(0).(model.ConvertedRequest), args.Error(1)
}
