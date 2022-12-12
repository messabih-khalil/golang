package main

import "fmt"

type Product struct {
	productName string
	inStock bool
}

type Store struct{
	products []Product
}

// receiver function
func (store * Store) rcvProductInStock(productId int){
	store.products[productId].inStock = true
}


func main() {
	store := Store{products: make([]Product, 5)}
	store.products[0] = Product{
		productName: "pro1",
		inStock: false,
	}
	store.rcvProductInStock(1)
	fmt.Println(store)
}
