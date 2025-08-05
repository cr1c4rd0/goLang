package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contactos
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guarda contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file) // convierte estructura de datos a formato JSON
	err = encoder.Encode(contacts)   // serializa los contactos a JSON
	if err != nil {
		return err
	}
	return nil
}

// Carga contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file) // decodifica el JSON a estructura de datos
	err = decoder.Decode(contacts)   // deserializa los contactos desde JSON
	if err != nil {
		return err
	}
	return nil
}

// Función principal
func main() {
	// Slice de contactos
	var contacts []Contact
	// Cargar contactos desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar contactos:", err)
	}
	// Crear instancia de bufio para crear contactos
	reader := bufio.NewReader(os.Stdin)
	for {
		// Motrar opciones al usuario
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Ver contactos")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer opción del usuario
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Opción inválida, intente de nuevo.")
			return
		}

		// Manejar opciones del usuario
		switch option {
		case 1: // Agregar contacto
			var c Contact
			fmt.Print("Ingrese nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Ingrese teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			// agregar contacto al slice
			contacts = append(contacts, c)
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar contacto:", err)
			} else {
				fmt.Println("Contacto guardado exitosamente.")
			}
		case 2: // Ver contactos
			fmt.Println("Contactos:")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s - Email: %s - Teléfono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
		case 3: // Salir
			fmt.Println("Saliendo...")
			return
		default: // Opción inválida
			fmt.Println("Opción inválida, intente de nuevo.")
			continue
		}
	}
}
