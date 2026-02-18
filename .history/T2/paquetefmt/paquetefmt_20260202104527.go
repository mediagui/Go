package paquetefmt

import (
	"fmt"
	"log"
)

func Formatfmt() {

	var intro string
	dato, err := fmt.Scanf("Dato:", intro)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Dato:", dato)

}
