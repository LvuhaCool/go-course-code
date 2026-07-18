package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance-usd < 0 {
		err := errors.New("Недостаточно средств!")
		return err
	}
	user.Balance -= usd
	return ""
}

func main() {
	user := User{
		Name:    "Олег",
		Balance: 10,
	}

	pp.Println("User до:", user)
	str := Pay(&user, 15)
	pp.Println("User после:", user)

	if str == "" {
		fmt.Println("Была произведена оплата!")
	} else {
		fmt.Println("Оплаты не было! Причина:", str)
	}
}
