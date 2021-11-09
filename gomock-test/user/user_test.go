package user

import (
	"gomock-test/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/prashantv/gostub"
)

func TestUse(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

}

func TestIncreaseAge(t *testing.T) {
	var configFile Config
	stubs := gostub.Stub(&configFile, "cnfg")
	defer stubs.Reset()

	oldAge := configFile.Age
	configFile.IncreaseAge()
	if oldAge+1 != configFile.Age {
		t.Error("age doesn't change")
	}
}
