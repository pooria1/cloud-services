package user

import (
	"gomock-test/doer"
)

type User struct {
	Doer doer.Doer
}

type Config struct {
	Name     string
	PastName string
	Age      int
}

func (u *User) Use() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}

func (c *Config) IncreaseAge() {
	c.Age++
}
