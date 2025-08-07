package main

import (
	"fmt"
	"log"

	"github.com/cr1c4rd0/goLang/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Cristian")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
