package main

import (
	"fmt"
	"slices"
)

func main() {
	var arregloCadenas []string	
	fmt.Println("datos:", "contenido", arregloCadenas, "condición:", arregloCadenas == nil, "longitud:", len(arregloCadenas) == 0)	

	arregloCadenas = make([]string, 3)	
	arregloCadenas[0] = "1"
	arregloCadenas[1] = "2"
	arregloCadenas[2] = "3"
	fmt.Println("datos:", arregloCadenas)
	fmt.Println("dato especifico:", arregloCadenas[2])
	fmt.Println("longitud:", len(arregloCadenas))
	
	arregloCadenas = append(arregloCadenas, "d")
	arregloCadenas = append(arregloCadenas, "e", "f")
	fmt.Println("valores finales:", arregloCadenas)
	fmt.Println("longitud:", len(arregloCadenas))

	segundoArreglo := []string{"g", "h", "i"}
	fmt.Println("dato 2:", segundoArreglo)

	tercerArreglo := []string{"g", "h", "i"}
	if slices.Equal(segundoArreglo, tercerArreglo) {
		fmt.Println("2 es igual a 3")
	} else {
		fmt.Println("no son iguales")
	}

	

}