package main

import (
	"fmt" 
	"unicode/utf8"
)


func main() {
	const saludo = "สวัสดี" // "Hola" en tailandés, contiene caracteres Unicode
	const saludo2 = "Hola"
	
	fmt.Println("El saludo es:", saludo)
	fmt.Println("Len:", len(saludo))

	// fmt.Println("El saludo2 es:", saludo2)
	// fmt.Println("Len:", len(saludo2))

	fmt.Println("\nRunas individuales Hola en Tailandes:")

	for i:= 0; i < len(saludo); i++ {
		fmt.Printf("%x ", saludo[i])
	}
	fmt.Println()
	fmt.Println("Conteo de runas:", utf8.RuneCountInString(saludo))
	fmt.Println()

	for idx, valorRuna := range saludo {
		fmt.Printf("%#U :comienza en %d\n", valorRuna, idx)
	}
}


