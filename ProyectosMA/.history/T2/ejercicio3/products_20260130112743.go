package products

import (
	"fmt"
	"math/rand"
	"time"
)

type productStruct struct {
	name  string
	prize float32
}

var (
	products []productStruct
	product  productStruct
)

func PrintProducts() {

	loadProducts()

	for i, v := range products {

		fmt.Println(i, "\t", v.name, "\t", v.prize, "\t€")

	}

}

func loadProducts() {

	for i := 1; i < 10; i++ {

		randomProduct := buildRandomProduct(i)

		products = append(products, randomProduct)

	}
}

func buildRandomProduct(i int) productStruct {

	var product productStruct

	product.name = fmt.Sprint("Producto ", i)
	product.prize = generateRandomCost(i)

	return product

}

func generateRandomCost(costSeed int) float32 {

	// Build the random generator using current time stamp
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	return generator.Float32() * float32(generator.Intn(costSeed))

}
