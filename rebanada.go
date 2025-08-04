package main

import "fmt"

func main() {
	diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}

	diaRebanada := diasSemana[0:5] // Rebanada de los primeros 5 días de la semana
	fmt.Println("Rebanada de la semana:", diaRebanada)

	// Agregar elementos a la rebanada
	diaRebanada = append(diaRebanada, "Nuevo Día")
	fmt.Println("Rebanada actualizada:", diaRebanada)

	fmt.Println(len(diaRebanada)) // Imprimir la longitud de la rebanada
	fmt.Println(cap(diaRebanada)) // Imprimir la capacidad de la rebanada

	nombre := make([]string, 5, 10) // Crear una rebanada con longitud 5 y capacidad 10
	nombre[0] = "Juan"
	fmt.Println("Rebanada de nombres:", nombre)

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	copy(rebanada2, rebanada1) // Copiar rebanada1 a rebanada2
	fmt.Println("Rebanada copiada:", rebanada2)
}
