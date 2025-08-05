package main

import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("Cristian")
	fmt.Println(message)
}
