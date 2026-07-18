package main

import (
	"errors"
	"fmt"
	"math/rand"

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
	return nil
}

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor-10 <= 0 {
		return 0, errors.New("Мы не стали газовать, чтобы не сломать машину")
	}

	kmph := rand.Intn(151)

	c.Armor -= 10

	return kmph, nil
}

func main() {

	defer func() {
		p := recover()

		if p != nil {
			fmt.Println("Была паника:", p)
		}
	}()

	slice := []int{1, 2, 3}
	fmt.Println(slice[4])

	user := User{
		Name:    "Олег",
		Balance: 10,
	}

	pp.Println("User до:", user)
	err := Pay(&user, 15)
	pp.Println("User после:", user)

	if err != nil {
		fmt.Println("Оплаты не было! Причина:", err.Error())
	} else {
		fmt.Println("Была произведена оплата!")
	}

	car := Car{
		Armor: 25,
	}

	for {
		pp.Println("car до:", car)
		kmph, err := car.Gas()
		pp.Println("car после:", car)

		if err != nil {
			fmt.Println("Ошибка нажатия на газ:", err.Error())
			break
		}
		fmt.Println("Получившийся разгон:", kmph)
		fmt.Println("")
	}
}
