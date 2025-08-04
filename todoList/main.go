package main

import (
	"fmt"
)

type Tarea struct {
	Nombre      string
	Descripcion string
	Completada  bool
}

type ListaDeTareas struct {
	Tareas []Tarea
}

// metodo para agregar una tarea a la lista
func agregarTarea(lista *ListaDeTareas, tarea Tarea) {
	lista.Tareas = append(lista.Tareas, tarea)
}

// Método para marcar como completada
func (lista *ListaDeTareas) marcarCompletada(indice int) {
	lista.Tareas[indice].Completada = true
}

// método para editar tarea
func (lista *ListaDeTareas) editarTarea(indice int, nuevaTarea Tarea) {
	lista.Tareas[indice] = nuevaTarea
}

// método para eliminar tarea
func (lista *ListaDeTareas) eliminarTarea(indice int) {
	lista.Tareas = append(lista.Tareas[:indice], lista.Tareas[indice+1:]...)
}

func main() {
	// Crear una nueva lista de tareas
	lista := ListaDeTareas{}

	for {
		var opcion int
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Marcar tarea como completada")
		fmt.Println("3. Editar tarea")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Mostrar lista de tareas")
		fmt.Println("6. Salir")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var nombre, descripcion string
			fmt.Println("Ingrese el nombre de la tarea:")
			fmt.Scanln(&nombre)
			fmt.Println("Ingrese la descripción de la tarea:")
			fmt.Scanln(&descripcion)
			tarea := Tarea{Nombre: nombre, Descripcion: descripcion, Completada: false}
			agregarTarea(&lista, tarea)
		case 2:
			var indice int
			fmt.Println("Ingrese el índice de la tarea a marcar como completada:")
			fmt.Scanln(&indice)
			lista.marcarCompletada(indice)
		case 3:
			var indice int
			var nombre, descripcion string
			fmt.Println("Ingrese el índice de la tarea a editar:")
			fmt.Scanln(&indice)
			fmt.Println("Ingrese el nuevo nombre de la tarea:")
			fmt.Scanln(&nombre)
			fmt.Println("Ingrese la nueva descripción de la tarea:")
			fmt.Scanln(&descripcion)
			nuevaTarea := Tarea{Nombre: nombre, Descripcion: descripcion, Completada: false}
			lista.editarTarea(indice, nuevaTarea)
		case 4:
			var indice int
			fmt.Println("Ingrese el índice de la tarea a eliminar:")
			fmt.Scanln(&indice)
			lista.eliminarTarea(indice)
		case 5:
			fmt.Println("Lista de tareas:")
			for i, tarea := range lista.Tareas {
				fmt.Printf("%d: %s - %s (Completada: %t)\n", i, tarea.Nombre, tarea.Descripcion, tarea.Completada)
			}
		case 6:
			fmt.Println("Saliendo del programa.")
			return
		default:
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}
}
