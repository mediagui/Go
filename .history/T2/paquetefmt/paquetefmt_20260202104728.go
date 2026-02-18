package paquetefmt

import (
	"fmt"
	"log"
)

func Formatfmt() {

	fmt.Print("Dato:")
	dato, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dato:", dato)

}
