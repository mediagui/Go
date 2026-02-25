package main

import (
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

	for i := 0; i < 1000; i++ {
		go readFromSlice(&slice, &mutex)
	}
	// readValue := readFromSlice(&slice, &mutex)
}

// Lee del slice recibido como parámetro y devuelve el valor de la posición random generada internamente
func readFromSlice(slice *[]string, m *sync.RWMutex) string {
	m.RLock()         // Bloquea para lectura
	defer m.RUnlock() //Desbloquea al final de la función

	randIndex := rand.Intn(len(*slice)) //Genera el id random
	return (*slice)[randIndex]          //Devuelve el valor
}

func writeToSlice(slice *[]string, value string, m *sync.RWMutex, wg *sync.WaitGroup) {

	m.Lock()         // Bloquea para escritura
	defer m.Unlock() //Desbloquea al final de la función

	*slice = append(*slice, value)
	wg.Done()
}
