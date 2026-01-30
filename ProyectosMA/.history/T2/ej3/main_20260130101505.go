package main

import "fmt"

type productStruct struct {
	nombre string
	precio float32
}

var productosSlice []productStruct
var producto productStruct

func main() {

	loadProducts()
	printProducts()

}

func printProducts() {
	for i, v := range productosSlice {

		fmt.Println(i, "\t", v.nombre, "\t", v.precio)

	}

}

func loadProducts() {

	for i := 0; i < 10; i++ {



		productosSlice = append(productosSlice, producto)

	}

	func buildProduct() productStruct{

		return productStruct{
			fmt.Sprint("Producto %f", i),
			float32(i) * 0.1
		}

	}

}
