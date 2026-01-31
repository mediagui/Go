package products

import (
	"ProyectosMA/T2/dto"
	"ProyectosMA/T2/format"
	"fmt"
	"math/rand"
	"time"
)

var (
	products []dto.ProductStruct
	product  dto.ProductStruct
)

func PrintProducts() {

	loadProducts()

	for i, v := range products {

		prize := format.PrintFormattedPrize(v.Prize)
		// prize := strconv.FormatFloat(float64(v.prize), 'f', 2, 32)

		fmt.Println(i, "\t", v.Name, "\t", prize, "€")

	}

}

func loadProducts() {

	for i := 1; i < 10; i++ {

		randomProduct := buildRandomProduct(i)

		products = append(products, randomProduct)

	}
}

func buildRandomProduct(i int) dto.ProductStruct {

	var product dto.ProductStruct

	product.Name = fmt.Sprint("Producto ", i)
	product.Prize = generateRandomPrize(i)

	return product

}

func generateRandomPrize(costSeed int) float32 {

	// Build the random generator using current time stamp
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	return generator.Float32() * float32(generator.Intn(costSeed))

}
