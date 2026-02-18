package main

import (
	"fmt"
	"math"
)

// Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

// Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.
// Calcular el área del triángulo usando la fórmula base x altura / 2.
// Calcular el perímetro del triángulo sumando los lados.

// Imprimir el área y el perímetro del triángulo con dos decimales de precisión.
// El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro.
// También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.
// Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.

// Ejemplo de entrada y salida:

// Ingrese lado 1: 3.5
// Ingrese lado 2: 4.2
// Área: 7.35
// Perímetro: 12.20

func main() {

	another := "Y"

	for another == "Y" || another == "y" {

		//Borra la consola
		fmt.Print("\033[H\033[2J")

		triangleCalculation()

		fmt.Print("\nOther calculation? [y/N]")
		fmt.Scanln(&another)

	}

}

func triangleCalculation() {
	side1, side2 := requestData()

	hipotenuse := hipotenuseCalculation(side1, side2)
	perimeter := perimeterCalculation(hipotenuse, side1, side2)
	area := areaCalculation(hipotenuse, side1)

	showData(area, perimeter, hipotenuse)
}

func areaCalculation(hipotenuse, height float64) float64 {
	return (hipotenuse * height) * 0.5
}

func requestData() (float64, float64) {

	var (
		cateto1 float64
		cateto2 float64
	)

	fmt.Print("Cateto 1: ")
	fmt.Scanln(&cateto1)

	fmt.Print("Cateto 2: ")
	fmt.Scanln(&cateto2)

	return cateto1, cateto2

}

func requestSide(cateto float64) float64 {

	fmt.Print("Cateto: ")
	fmt.Scanln(&cateto)
	return cateto
}

func hipotenuseCalculation(side1, side2 float64) float64 {
	return math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2))
}

func perimeterCalculation(hipotenuse, side1, side2 float64) float64 {
	return hipotenuse + side1 + side2
}

func showData(area, perimeter, hipotenuse float64) {

	const PRECISION int = 2

	fmt.Println("Datos:")

	fmt.Printf("\tArea: %.*f\n", PRECISION, area)
	fmt.Printf("\tHipotenusa: %.*f\n", PRECISION, hipotenuse)
	fmt.Printf("\tPerímetro: %.*f\n", PRECISION, perimeter)
}
