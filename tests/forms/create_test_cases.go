package forms

import (
	"Assessment/model"
	"Assessment/tests"
	"github.com/stretchr/testify/mock"
)

var FormCreateCases = []FormTestCase{
	{
		Case:  "Valid Case with proper value",
		Input: CreateInput,
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				mockCall := tests.MockFormsRepo.On("Create", CreateInput).Return(
					CreateOutput, nil)
				return mockCall
			},
		},
		ExpectedOutPut: CreateOutput,
		ExpectedError:  nil,
	},
	{
		Case:  "Valid Case without value",
		Input: map[string]string{},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				mockCall := tests.MockFormsRepo.On("Create", map[string]string{}).Return(
					model.ConvertedRequest{}, OutputError)
				return mockCall
			},
		},
		ExpectedOutPut: model.ConvertedRequest{},
		ExpectedError:  OutputError,
	},
}
