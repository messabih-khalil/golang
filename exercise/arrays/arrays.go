//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

// product object
type Product struct{
	name string
	price float32
}

func arrayLength(products[4] Product) int{
	length := len(products)
	return length
}

func lastItem(products[4] Product) Product{
	lastEL := products[2]
	return lastEL
}


func productsCost(products[4] Product) int {
	var total int
	for i := 0; i < len(products); i++ {
		product := products[i]
		total += int(product.price)
	}

	return total
}


func main() {
	// products
	products := [4]Product{
		{"ps4" , 300},
		{"ps5" , 700},
		{"xbox series s" , 400},
	}

	lastEl := lastItem(products)
	
	fmt.Println("last element is :" , lastEl.name , "and the price is :" , lastEl.price)

	// total number of products
	totalProducts := arrayLength(products)
	fmt.Println("the number of products is :" , totalProducts)

	// total cost 
	totalCost := productsCost(products)
	fmt.Println("the total price of products is :" , totalCost)

	// add new product

	newProduct := Product{"alienware" , 2000}

	products[3] = newProduct
	
	fmt.Println("the new elements is :" , products[3].name , "and the price is :" , products[3].price)
}
