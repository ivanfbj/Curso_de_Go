package main

import "fmt"

func suma(entero1 int, entero2 int) int {
	result := entero1 + entero2
	return result
}

func sumaLarga(entero1, entero2, entero3, entero4 int) int {
	result := entero1 + entero2 + entero3 + entero4
	return result
}

func main() {
	var numero1, numero2, numero3, numero4 int

	fmt.Println("Ingresa el primer numero: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingresa el segundo numero: ")
	fmt.Scanln(&numero2)
	fmt.Println("Ingresa el tercer numero: ")
	fmt.Scanln(&numero3)
	fmt.Println("Ingresa el cuarto numero: ")
	fmt.Scanln(&numero4)

	fmt.Println("--------------------------------")

	fmt.Println("Resultado de la suma: ", suma(numero1, numero2))
	fmt.Println("Resultado de la suma larga: ", sumaLarga(numero1, numero2, numero3, numero4))
}