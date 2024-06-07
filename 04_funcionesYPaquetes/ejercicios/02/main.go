package main

import (
	"fmt"
)

func main() {
	doble, triple := doubleAndTriple(2)
	fmt.Println(doble, triple)
}

func doubleAndTriple(value int) (doubled, tripled int) {
	doubled = value * 2
	tripled = value * 3
	return
}
