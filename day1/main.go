package main

import "fmt"

type payment struct{}
type razorpay struct{}
type stripe struct{}

func (r razorpay) makepayment(amount float64) {
	fmt.Println("Payment made using Razorpay", amount)
}

func (s stripe) makepayment(amount float64) {
	fmt.Println("Payment made using Stripe", amount)
}

func (p payment) makepayment(amount float64) {
	payment := stripe{}
	payment.makepayment(amount)
}

func main() {
	payment := payment{}
	payment.makepayment(1000.10)
}
