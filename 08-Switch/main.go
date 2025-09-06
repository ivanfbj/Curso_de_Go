package main

import (
	"fmt"
	"time"
)

func main() {
	
	i := 2
	fmt.Println("Escribe", i, "como")


	switch i {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("A estudiar!!")
	default:
		fmt.Println("A trabajar!!")
	}

	tiempo := time.Now()
	fmt.Println("La hora es: ", tiempo)

	switch {
	case tiempo.Hour() < 12:
		fmt.Println("Buenos días!!")
	default:
		fmt.Println("Buenas tardes!!")
	}

	dia := time.Now().Weekday()

	switch dia {
	case time.Monday:
		fmt.Println("Hoy es lunes")
	case time.Tuesday:
		fmt.Println("Hoy es martes")
	case time.Wednesday:
		fmt.Println("Hoy es miércoles")
	case time.Thursday:
		fmt.Println("Hoy es jueves")
	case time.Friday:
		fmt.Println("Hoy es viernes")
	case time.Saturday:
		fmt.Println("Hoy es sábado")
	case time.Sunday:
		fmt.Println("Hoy es domingo")
	}
	
	
}