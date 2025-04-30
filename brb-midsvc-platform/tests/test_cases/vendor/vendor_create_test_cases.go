package vendor

import (
	"Assessment/model"
	"Assessment/tests"
	"errors"

	"github.com/stretchr/testify/mock"
)

type VendorCreateTestCase struct {
	Case          string
	Vendor        model.Vendor
	MockFunctions []func() *mock.Call
	ExpectedError error
}

var VendorCreateTestCases = []VendorCreateTestCase{
	{
		Case:   "Valid Case",
		Vendor: model.Vendor{
			Name:        "Test Vendor",
			Services: []model.Service{},
		},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				mockCall := tests.MockVendorRepo.On("Create",  &model.Vendor{
					Name:        "Test Vendor",
					Services: []model.Service{},
				}).Return(nil)
				return mockCall
			},
		},
		ExpectedError: nil,
	},
	{
		Case:   "Error Scenario",
		Vendor: model.Vendor{},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				mockCall := tests.MockVendorRepo.On("Create",  &model.Vendor{}).Return(errors.New("created failed"))

				return mockCall
			},
		},
		ExpectedError: errors.New("created failed"),
	},
}