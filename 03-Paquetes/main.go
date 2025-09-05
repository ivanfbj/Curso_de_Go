package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("La variable HOME no está configurada")
	} else {
		fmt.Println("HOME sí está configurada y es:", envVar)
	}

	file, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Println("Error en creación del archivo:", err)
		return
	}
	defer file.Close()
	fmt.Println("Archivo creado exitosamente")
}