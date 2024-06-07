package main

import (
	"fmt"
)

func main() {
	var a float64 = 7.5
	var b int = 3

	aConvertido := int(a)
	bConvertido := float64(b)
	sumFloat := float64(a + bConvertido)
	sumInt := int(aConvertido + b)

	// otra resolución
	//var sumFloat float64 = a + float64(b)
	//var sumInt int = int(a) + b

	fmt.Printf("Suma como float64: %.2f\n", sumFloat)
	fmt.Printf("Suma como int: %d\n", sumInt)
}
