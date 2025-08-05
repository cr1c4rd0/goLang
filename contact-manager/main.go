package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("división por cero no es permitida")
	}
	return dividendo / divisor, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error al dividir:", err)
		return
	}
	fmt.Println("Resultado de la división:", result)

	str := "123f"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error al convertir de string a entero:", err)
		return
	}
	fmt.Println("Convertido a entero:", num)
}
