package main

import "fmt"

// PaymentStrategy is a small interface representing a payment strategy.
type PaymentStrategy interface {
	Pay(amount float64)
}

// CreditCardPayment is an implementation of the PaymentStrategy for credit card payments.
type CreditCardPayment struct{}

func (cc *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card\n", amount)
}

// CashPayment is an implementation of the PaymentStrategy for cash payments.
type CashPayment struct{}

func (cash *CashPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f in Cash\n", amount)
}

// ShoppingCart represents a consumer that uses a PaymentStrategy.
type ShoppingCart struct {
	paymentMethod PaymentStrategy
}

// SetPaymentMethod sets the payment strategy for the shopping cart.
func (cart *ShoppingCart) SetPaymentMethod(paymentMethod PaymentStrategy) {
	cart.paymentMethod = paymentMethod
}

// Checkout processes the payment using the selected strategy.
func (cart *ShoppingCart) Checkout(amount float64) {
	if cart.paymentMethod == nil {
		fmt.Println("Payment method not set.")
		return
	}
	cart.paymentMethod.Pay(amount)
}

func main() {
	// Client code using the strategy pattern
	creditCardPayment := &CreditCardPayment{}
	cashPayment := &CashPayment{}

	shoppingCart := &ShoppingCart{}

	// Use credit card payment strategy
	shoppingCart.SetPaymentMethod(creditCardPayment)
	shoppingCart.Checkout(100.0)

	// Use cash payment strategy
	shoppingCart.SetPaymentMethod(cashPayment)
	shoppingCart.Checkout(50.0)
}
