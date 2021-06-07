package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
func main() {
	category := [4]string{"fashion", "electronics", "sport", "food"}
	products := [20]Product{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}

	for _, product := range products {
		fmt.Println(product)
	}

	fmt.Println("------Tìm sản phẩm theo tên-------")
	var nameProduct string
	fmt.Print("Enter name of product: ")
	fmt.Scan(&nameProduct)
	for _, product := range products {
		if strings.Contains(product.Name, nameProduct) {
			fmt.Println(product)
		}
	}

	fmt.Println("------Tìm sản phẩm theo category-------")
	var cat string
	fmt.Print("Enter the category: ")
	fmt.Scan(&cat)
	for _, product := range products {
		if strings.Compare(product.Category, cat) == 0 {
			fmt.Println(product)
		}
	}

	fmt.Println("------Tìm sản phẩm theo giá-------")
	var max, min int
	fmt.Print("Enter max of price: ")
	fmt.Scan(&max)
	fmt.Print("Enter min of price: ")
	fmt.Scan(&min)
	for _, product := range products {
		if product.Price >= min && product.Price <= max {
			fmt.Println(product)
		}
	}
}
