package main

import "fmt"

type Customer struct {
	name     string
	fidelity int
}

type CartItem struct {
	product  string
	quantity int
	price    float64
}

func (item CartItem) Total() float64 {

	return item.price * float64(item.quantity)

}

func (item CartItem) String() {
	fmt.Sprintf("product-name : %s product-quantity : %d product-price : %2f", item.product, item.quantity, item.price)
}

// order is the relatinship a customer , the cart and promo

type Order struct {
	ctm   Customer
	cart  []CartItem
	promo Promotion
}

// Total is the sum of items purchased

func (order Order) Total() float64 {
	total := 0.0
	for _, item := range order.cart {

		total += item.Total()
	}

	return total

}

// Due calculates order value considering discount

func (order Order) Due() float64 {
	discount := 0.0

	if order.promo != nil {
		discount = order.promo.Discount(order)
	}

	return order.Total() - discount
}

// String represents the order when its printed

func (order Order) String() {
	fmt.Sprintf("order total : %2f due : %2f", order.Total(), order.Due())
}

type Promotion interface {
	Discount(Order) float64
}

// FidelityPromo is a conrete implementation of the Promotion
type FidelityPromo struct{}

func (FidelityPromo) Discount(order Order) float64 {
	if order.ctm.fidelity >= 10000 {
		return order.Total() * 0.05
	}
	return 0.0
}
