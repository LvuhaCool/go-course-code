package main

import (
	"github.com/k0kubun/pp"
)

type User struct {
	Name string
	Age  int
}

func (u User) GetAge() int {
	return u.Age
}

func main() {
	user1 := User{
		Name: "LvuhaCool",
		Age:  13,
	}

	pp.Println(user1.GetAge())

	pp.Println(user1)
}
