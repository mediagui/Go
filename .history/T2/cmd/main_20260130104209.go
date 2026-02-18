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
