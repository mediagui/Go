package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1000)

	slice := make([]string, 0)
	mutex := sync.RWMutex{}

	for i := 0; i < 1000; i++ {
		go writeToSlice(&slice, "Valor "+string(rune(i+65)), &mutex, &wg)
	}

	wg.Wait() // Espera a que todas las goroutines terminen

	// Canal para recibir los valores
	resultChannel := make(chan string, 1000)

	for range 1000 {
		go readFromSlice(&slice, &mutex, resultChannel)
	}

	// Leer e imprimir los valores del canal
	for range 1000 {
		value := <-resultChannel
		fmt.Println(value)
	}
}

// Lee del slice recibido como parámetro y envía el valor al canal
func readFromSlice(slice *[]string, m *sync.RWMutex, ch chan string) {
	m.RLock()         // Bloquea para lectura
	defer m.RUnlock() //Desbloquea al final de la función

	log.Println("Leyendo del slice...")

	randIndex := rand.Intn(len(*slice)) //Genera el id random
	value := (*slice)[randIndex]        //Obtiene el valor
	ch <- value                         //Envía el valor al canal
}

func writeToSlice(slice *[]string, value string, m *sync.RWMutex, wg *sync.WaitGroup) {

	m.Lock()         // Bloquea para escritura
	defer m.Unlock() //Desbloquea al final de la función

	log.Println("Escribiendo en el slice...")
	*slice = append(*slice, value)
	wg.Done()
}
