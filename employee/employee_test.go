package employee_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/unit-test-example/employee"
	"github.com/unit-test-example/test/mocks"
	"strconv"
	"testing"
)

func TestGetData(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRedisCache := mocks.NewMockSample(mockCtrl)

	emp := &employee.Employee{RedisCache: mockRedisCache}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	//mockRedisCache.EXPECT().DoSomething(123, "Hello Babulal").Return("test abc", nil).Times(1)
	//mockRedisCache.EXPECT().DoSomething(123, "Hello test").Return("test abc", nil).Times(1)
	mockRedisCache.EXPECT().DoSomething(gomock.Any(), gomock.Any()).DoAndReturn(func(id int, name string) (string, error) {

		data := fmt.Sprintf("Test : %s / %s / %s", strconv.Itoa(id), name, "pune")
		return data, nil
	})

	resp, err := emp.GetData()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
	assert.NotEmpty(t, resp)

}
