package main

type store struct {
}

type product struct {
	id    int
	name  string
	price float64
}

type Cart struct {
	id      int
	name    string
	product []product
}
