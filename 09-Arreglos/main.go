package main

import "fmt"

func main() {
	var arreglo [5]int
	fmt.Println("Arreglo completo:", arreglo)

	arreglo[4] = 100
	fmt.Println("Arreglo completo:", arreglo)
	fmt.Println("Elemento en la posición cuatro:", arreglo[4])

	fmt.Println("Tamaño del arreglo:", len(arreglo))

	lista := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Lista original:", lista)

	lista2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("Tamaño del arreglo inferido:", len(lista2))
}
