package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
	fmt.Println("¡Juego terminado! Gracias por jugar.")
}

func jugar() {
	numAleatorio := rand.Intn(100) + 1 // para que nunca me salga 0
	var numIngresasdo int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un número entre 1 y 100 (intentos restantes: %d): ", maxIntentos-(intentos+1))
		fmt.Scanln(&numIngresasdo)

		if numIngresasdo == numAleatorio {
			fmt.Println("¡Felicidades! Adivinaste el número.")
			return
		} else if numIngresasdo < numAleatorio {
			fmt.Println("El número es mayor. Intenta de nuevo.")
		} else {
			fmt.Println("El número es menor. Intenta de nuevo.")
		}
	}
	fmt.Printf("Lo siento, has agotado tus intentos. El número era %d.\n", numAleatorio)
}
