package main

import "fmt"

type paymentmethod interface {
	pay(amount float64)
}

type payment struct {
	paymentmethod
}
type razorpay struct{}
type stripe struct{}
type juspay struct{}

func (r razorpay) pay(amount float64) {
	fmt.Println("Payment made using Razorpay", amount)
}

func (s stripe) pay(amount float64) {
	fmt.Println("Payment made using Stripe", amount)
}

func (s juspay) pay(amount float64) {
	fmt.Println("Payment made using juspay", amount)
}

func (p payment) makepayment(amount float64) {
	p.pay(amount)
}

func main() {
	p := payment{
		paymentmethod: razorpay{},
	}
	p.makepayment(1000.10)
}
