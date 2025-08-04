package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":     "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
		"amarillo": "#FFFF00",
	}
	fmt.Println(colors["rojo"])
	colors["negro"] = "#000000"
	fmt.Println(colors)
	valor, ok := colors["rojo"]
	fmt.Println("el Valor del rojo es: ", valor, ok)
	if ok {
		fmt.Println("El color rojo está en el mapa.")
	} else {
		fmt.Println("El color rojo no está en el mapa.")
	}

	// Eliminacion de elementos
	delete(colors, "rojo")
	fmt.Println("Después de eliminar el rojo:", colors)
}
