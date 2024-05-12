package main

type Item struct {
	id               int
	price            float64
	products         []product
	numberOfProducts int
}

func (item *Item) addProducts(Product product) {
	item.products = append(item.products, Product)
	item.numberOfProducts++
}

func (item *Item) removeProducts(Product int) {
	for product := 0; product < item.numberOfProducts; product++ {
		if item.products[product].id == Product {
			item.products = append(item.products[:product], item.products[product+1:]...)
			item.numberOfProducts--
		}
	}
	return
}

func (item *Item) setId(id int) {
	item.id = id
}

func (item *Item) getId() int {
	return item.id
}

func (item *Item) getNumberOfProducts() int {
	return item.numberOfProducts
}

func (item *Item) getPrice() float64 {
	totalPrice := 0.0
	for _, product := range item.products {
		totalPrice += product.price
	}
	return totalPrice
}
