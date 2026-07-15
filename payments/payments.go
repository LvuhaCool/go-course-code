package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentMethod PaymentMethod
}

func NewPaymentModule(PaymentMethod PaymentMethod) PaymentModule {
	return PaymentModule{
		paymentMethod: paymentMethod(),
	}
}

func (p PaymentModule) Pay(description string, usd int) int {}

func (p PaymentModule) Cancel() {}

func (p PaymentModule) Info() {}

func (p PaymentModule) AllInfo() {}
