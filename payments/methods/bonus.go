package methods

import (
	"fmt"
	"math/rand"
)

type Bonus struct{}

func NewBonus() Bonus {
	return Bonus{}
}

func (c Bonus) Pay(usd int) int {
	fmt.Println("Оплата бонусами!")
	fmt.Println("Размер оплаты:", usd*100, "бонусов")

	id := rand.Int()

	return id
}

func (c Bonus) Cancel(id int) {
	fmt.Println("Отмена операции по оплате бонусами! ID:", id)
}
