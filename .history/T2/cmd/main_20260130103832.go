package main

import (
	"fmt"
	"math/rand"
	"time"
)

type productStruct struct {
	nombre string
	precio float32
}

var products []productStruct
var producto productStruct

func main() {

	loadProducts()
	printProducts()

}

func printProducts() {
	for i, v := range products {

		fmt.Println(i, "\t", v.nombre, "\t", v.precio)

	}

}

func loadProducts() {

	for i := 0; i < 10; i++ {

		randomProduct := buildRandomProduct(i)

		products = append(products, randomProduct)

	}
}

func buildRandomProduct(i int) productStruct {

	var product productStruct

	product.nombre = fmt.Sprint("Producto ", i)
	product.precio = generateRandomCost(i)

	return product

}

func generateRandomCost(costSeed int) float32 {

	// Build the random generator using current time stamp
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	return generator.Float32() * float32(generator.Intn(costSeed*100))

}
