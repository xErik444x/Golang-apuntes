package main

import "fmt"

func main() {
	switch number := 10; {
	case number < 0:
		fmt.Println("El número es negativo")
	case number == 0:
		fmt.Println("El número es cero")
	}
}
