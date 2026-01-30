package main

import "fmt"

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

	product.nombre = fmt.Sprint("Producto %f", i)
	product.precio = float32(i) * 0.1

	return product

}
