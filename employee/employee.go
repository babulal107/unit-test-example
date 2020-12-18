package employee

import (
	"fmt"
	"github.com/unit-test-example/cache"
)

type Employee struct {
	RedisCache cache.Sample
}

func (u *Employee) GetData() (data string, err error) {
	data, err = u.RedisCache.DoSomething(123, "Hello Babulal")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	return
}
