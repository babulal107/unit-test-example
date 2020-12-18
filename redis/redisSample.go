package redis

import (
	"fmt"
	"strconv"
)

type SampleRedisCache struct {

}

func (s *SampleRedisCache) DoSomething(id int, name string) (string, error) {
	str := fmt.Sprintf("User : %s / %s / %s", strconv.Itoa(id), name, "Pune")
	return str, nil
}
