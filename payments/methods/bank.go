package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Банковская оплата!")
	fmt.Println("Размер оплаты:", usd, "долларов")

	id := rand.Int()

	return id
}

func (c Bank) Cancel(id int) {
	fmt.Println("Отмена банковской операции! ID:", id)
}
