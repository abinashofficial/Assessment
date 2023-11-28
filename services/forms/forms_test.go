package forms

import (
	"Assessment/model"
	mongo_serv "Assessment/store"
	"Assessment/tapcontext"
	"Assessment/tests"
	"Assessment/tests/forms"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var mockRepos mongo_serv.Store
var TestFormService Service

func TestMain(m *testing.M) {
	mockRepos = mongo_serv.Store{
		FormsRepo: tests.MockFormsRepo,
	}

	TestFormService = New(mockRepos)

	fmt.Println("Starting Forms Service Relates Test Cases")
	code := m.Run()
	fmt.Println("Done with Forms Service Relates Test Cases")

	os.Exit(code)
}

func TestCreateForm(t *testing.T) {
	var ctx = tapcontext.NewTapContext()
	ctx.Locale = "en"
	for _, test := range forms.FormCreateCases {
		t.Run(test.Case, func(t *testing.T) {
			mockCalls := tests.InitializeMockFunctions(test.MockFunctions)
			form, err := TestFormService.Create(test.Input)
			if test.ExpectedError == nil {
				// Positive tests case
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedOutPut, form)
			} else {
				// Error scenario
				assert.Equal(t, "empty map", err.Error())
				assert.Equal(t, model.ConvertedRequest{}, form)
				assert.Error(t, err)
			}
			tests.UnsetMockFunctionCalls(mockCalls, []string{})
		})
	}
}
