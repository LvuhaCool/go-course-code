package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCrypto()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Бургер", 5)
	paymentModule.Pay("Телефон", 500)
	paymentModule.Pay("Игра", 20)

	allInfo := paymentModule.AllInfo()

	pp.Println("Все наши оплаты:", allInfo)
}
