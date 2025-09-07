package main

import "fmt"

func suma(numeros ...int) {
	fmt.Println("--------------------------------")
	fmt.Println(numeros, " <<--numeros")
	suma := 0
	for _, numero := range numeros {
		suma += numero
	}
	fmt.Println("La suma es: ", suma)
}

func main() {
	suma(1, 2)
	suma(1, 2, 3)

	numeros := []int{1, 2, 3, 4, 5}
	suma(numeros...)
}