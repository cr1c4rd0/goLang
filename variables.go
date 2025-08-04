package main

import (
	"fmt"
)

// Constantes
const pi float64 = 3.14159

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("¡Hola, mundo!")

	//Declaración de variables
	var firtName string = "Cristian"
	fmt.Println("Nombre:", firtName)
	var age int = 37
	fmt.Println("Edad:", age)

	fmt.Println("Valor de pi:", pi)
	fmt.Println("Día de la semana:", Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)

	//Imprimir caracteres especiales
	fullname := "Cristian \t(alias \"Cris\")\n"
	fmt.Println("Nombre completo:", fullname)

	//valores por defecto de variables
	// En Go, las variables no inicializadas tienen valores por defecto.
	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defauultBool  bool
		defaultString string
	)
	fmt.Println("Valores por defecto:")
	fmt.Println("int:", defaultInt)
	fmt.Println("uint:", defaultUint)
	fmt.Println("float32:", defaultFloat)
	fmt.Println("bool:", defauultBool)
	fmt.Println("string:", defaultString)
}
