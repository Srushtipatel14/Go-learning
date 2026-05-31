package main

import "fmt"

type OrderStatus string

const (
	shipped   OrderStatus = "shipped"
	delivered OrderStatus = "delivered"
	returned  OrderStatus = "returned"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status", status)
}

func main() {
	changeOrderStatus(delivered)
}
