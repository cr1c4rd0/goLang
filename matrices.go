package main

import "fmt"

func main() {
	var matriz = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(matriz); i++ {
		fmt.Printf("Elemento %d: %d\n", i, matriz[i])
	}

	//matriz bidimensional
	var matriz2 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Matriz bidimensional:", matriz2)
}
