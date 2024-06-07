package main

import "fmt"

// - Crea un `slice` de `enteros` con al menos `cinco elementos`.
// `Agrega dos elementos` más al slice y
// `muestra la longitud y capacidad`
// del slice después de agregar los nuevos elementos.
func main() {
	slice := make([]int, 5, 8)
	slice = append(slice, 1, 2)
	fmt.Printf("longitud: %d", len(slice))
	fmt.Println()
	fmt.Printf("capacidad: %d", cap(slice))
}
