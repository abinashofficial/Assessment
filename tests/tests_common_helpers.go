package tests

import (
	"Assessment/tests/mock_repos"
	"Assessment/utils"
	"github.com/stretchr/testify/mock"
)

var MockFormsRepo = new(mock_repos.MockFormRepo)

func UnsetMockFunctionCalls(mockCalls []*mock.Call, skipMockCallUnsetMethods []string) {
	mockCallsList := make([]string, 0)
	for _, mockCall := range mockCalls {
		if !utils.CheckKeyInSlice(skipMockCallUnsetMethods, mockCall.Method) && !utils.CheckKeyInSlice(mockCallsList, mockCall.Method) {
			mockCallsList = append(mockCallsList, mockCall.Method)
			mockCall.Unset()
		}
	}
}

func InitializeMockFunctions(functions []func() *mock.Call) []*mock.Call {
	var mockCalls = make([]*mock.Call, 0)
	for _, function := range functions {
		mockCall := function()
		mockCalls = append(mockCalls, mockCall)
	}
	return mockCalls
}
