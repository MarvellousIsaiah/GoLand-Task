package main



type customer struct {
	id      int
	cart    Cart
	name    string
	balance float64
}

func (customer *customer) getCart() Cart {
	return customer.cart

}
func (customer *customer) getBalance(cart Cart) float64 {
	    return 0
}

func