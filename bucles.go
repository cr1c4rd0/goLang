package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("Buenos días!")
	} else if hora < 18 {
		fmt.Println("Buenas tardes!")
	} else {
		fmt.Println("Buenas noches!")
	}

	// Otra forma de hacerlo declarando la variable en el if
	if t2 := time.Now(); t2.Hour() < 12 {
		fmt.Println("Buenos días!")
	} else if t2.Hour() < 18 {
		fmt.Println("Buenas tardes!")
	} else {
		fmt.Println("Buenas noches!")
	}

	// Ejemplo de switch
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("go run -> windows")
	case "linux":
		fmt.Println("go run -> linux")
	case "darwin":
		fmt.Println("go run -> darmacOSwin")
	default:
		fmt.Println("go run -> otro sistema operativo")
	}

	// Ejemplo de for

	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Otra forma de hacerlo asignado la variable en el for
	for j := 1; j <= 10; j++ {
		if j == 5 {
			break
		}
		fmt.Println(j)
	}

	for j := 1; j <= 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}
}
