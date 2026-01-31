package main

type producto struct {
	nombre string
	precio float32
}

var productos []producto

func main() {

}

func loadProducts() {

	for i := 0; i < 10; i++ {

		producto.nombre = "Producto" + i
		producto.precio = float(i)
	}

}
