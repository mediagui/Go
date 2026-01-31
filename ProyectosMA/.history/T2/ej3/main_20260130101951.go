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

		randomProduct := buildRandomProduct()

		products = append(products, randomProduct)

	}
}

func buildRandomProduct() productStruct {

		var product productoSproductStruct

		produto.nombre = fmt.Sprint("Producto %f", i)
		producto.precio = float32(i) * 0.1

		return producto

	}

}
