package main

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
	}
}

func loadProducts() {

	for i := 0; i < 10; i++ {
		producto.nombre = "Producto" + string(i)
		producto.precio = float32(i)

		productosSlice = append(productosSlice, producto)

	}

}
