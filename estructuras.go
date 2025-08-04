package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

func main() {
	var p Persona
	p.Nombre = "Juan"
	p.Edad = 30
	p.Email = "juan@example.com"

	fmt.Println(p)

	// otra forma de inicializar
	p2 := Persona{"Cristian", 37, "cristian@example.com"}
	fmt.Println(p2)
}
