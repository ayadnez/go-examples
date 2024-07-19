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

type BulkItemPromo struct{}

func (b BulkItemPromo) Discount(o Order) float64 {
	discount := 0.0

	for _, item := range o.cart {

		if item.quantity >= 20 {
			discount += item.Total() * 0.1

		}
	}
	return discount

}

type LargeOrderItem struct{}

func (l LargeOrderItem) Discount(o Order) float64 {

	set := map[string]bool{}

	for _, item := range o.cart {
		set[item.product] = true
	}

	if len(set) >= 10 {
		return o.Total() * 0.07
	}

	return 0.0
}

func main() {

	joe := Customer{"john doe", 0}
	ann := Customer{"anne bella ", 11000}

	cart := []CartItem{
		CartItem{"banana", 4, 0.50},
		CartItem{"apple", 5, 1.50},
		CartItem{"watermelon", 5, 4.00},
	}

	fmt.Printf("\n %s have %d fidelity points \n", joe.name, joe.fidelity)

	fmt.Println(Order{joe, cart, FidelityPromo{}})
	//fmt.Println(ann)
	fmt.Printf("\n %s have %d fidelity points \n", ann.name, ann.fidelity)
	fmt.Println(Order{ann, cart, FidelityPromo{}})

	bananaCart := []CartItem{
		CartItem{"banana", 30, 1.5},
		CartItem{"apple", 20, 2.0},
	}

	fmt.Printf("\n %s buy many items of the same product %s \n", joe.name, bananaCart)
	fmt.Println(Order{joe, bananaCart, BulkItemPromo{}})

	largeOrder := []CartItem{}

	for i := 0; i < 10; i++ {

		largeOrder = append(largeOrder, CartItem{string(65 + i), 1, 1.0})
	}

	fmt.Printf("\n %s representd orders with many discount items %s\n", joe.name, largeOrder)

	fmt.Println(Order{joe, largeOrder, LargeOrderItem{}})

}
