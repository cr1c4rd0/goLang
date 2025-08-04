package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.Nombre)
}

func main() {
	p := Persona{"Juan", 30, "juan@example.com"}
	p.diHola()

	var x int = 10
	fmt.Println("Valor de x:", x)
	editar(&x)
	fmt.Println("Valor de x después de editar:", x)
}

func editar(x *int) {
	*x = 20
}
