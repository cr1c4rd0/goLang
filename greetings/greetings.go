package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Varios saluos
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// enviar saludos aleatorios
func randomFormat() string {
	formats := []string{
		"Hello, %v. Welcome!",
		"Que bueno verte, %v",
		"Saludos, %v Encantado de conocerte",
	}
	return formats[rand.Intn(len(formats))]
}
