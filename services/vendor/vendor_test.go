package vendor

import (
	sqlStore "Assessment/store"
	"Assessment/tests"
	"Assessment/tests/test_cases/vendor"
	"fmt"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

var mockRepos sqlStore.Store
var TestVendorServ VendorService

func TestMain(m *testing.M) {	
	mockRepos = sqlStore.Store{
		VendorRepo: tests.MockVendorRepo,
	}
	TestVendorServ = NewVendorService(mockRepos.VendorRepo)


	fmt.Println("Starting Vendor Related Test Cases")
	code := m.Run()
	fmt.Println("Done with Vendor Service Relates Test Cases")

	os.Exit(code)
}

func TestCreateVendor(t *testing.T) {

	for _, tc := range vendor.VendorCreateTestCases {
		t.Run(tc.Case, func(t *testing.T) {
			mockCalls := tests.InitializeMockFunctions(tc.MockFunctions)
			err := TestVendorServ.CreateVendor(&tc.Vendor)
			if tc.ExpectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.ExpectedError.Error(), err.Error())
			}
			tests.UnsetMockFunctionCalls(mockCalls, []string{})
		})
	}
}
