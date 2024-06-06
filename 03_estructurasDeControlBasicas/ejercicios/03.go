package main

import "fmt"

func main() {
	day := "lunes"

	switch day {
	case "lunes":
		fmt.Println("Es lunes")
	case "martes":
		fmt.Println("Es martes")
	default:
		fmt.Println("No es ni lunes ni martes")
	}
}
