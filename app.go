package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"bufio"
	"os"
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
	fmt.Println("All products")
	for _, product := range products {
		fmt.Println(product)
	}
	productsSlice := products[:]
	sort.Slice(productsSlice, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
	fmt.Println("Top 5 price")
	for i := 0; i < 5; i++ {
		fmt.Println(productsSlice[i])
	}
	fmt.Println("--Đếm sản phẩm theo category----")
	categoryCount := map[string]int{}
	for _, product := range products {
		if _, ok := categoryCount[product.Category]; ok {
			categoryCount[product.Category]++
		} else {
			categoryCount[product.Category] = 0
		}
	}
	for key, value := range categoryCount {
		fmt.Println(key, value)
	}
	fmt.Println("------------------")
	 
}
func returnName(name string,products ...Product) {
	reader := bufio.NewReader(os.Stdin)
	for 
}
