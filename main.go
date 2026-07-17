package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewBonus()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Бургер", 5)
	idPhone := paymentModule.Pay("Телефон", 500)
	idGame := paymentModule.Pay("Игра", 20)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()

	pp.Println("Все наши оплаты:", allInfo)

	gameInfo := paymentModule.Info(idGame)
	pp.Println("Информация об игре:", gameInfo)
}
