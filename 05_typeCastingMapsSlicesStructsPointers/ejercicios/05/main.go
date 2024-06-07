package main

import (
	"fmt"
)

func duplicarValor(p *int) {
	*p *= 2
}

func main() {
	var x int = 10
	fmt.Println("Antes de duplicar:", x)
	duplicarValor(&x)
	fmt.Println("Después de duplicar:", x)
}
