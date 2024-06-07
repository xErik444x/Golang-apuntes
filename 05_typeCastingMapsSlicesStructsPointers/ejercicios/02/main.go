package main

import "fmt"

func main() {
	var clase map[string]int = map[string]int{
		"Alberto": 5,
		"Jose":    8,
		"Manuel":  2,
	}
	for nombre, nota := range clase {
		fmt.Println(nombre, nota)
	}
}
