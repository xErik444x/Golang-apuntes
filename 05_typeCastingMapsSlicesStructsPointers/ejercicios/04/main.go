package main

import "fmt"

/*
Define una `struct` llamada `Libro` con los campos `Titulo` y `Autor`.
Crea una instancia de Libro, inicializa sus campos y muestra sus valores.
*/
func main() {
	type Libro struct {
		Titulo string
		Autor  string
	}
	var libro Libro = Libro{"Las aventuras de Chowder", "juanito el tercero"}
	fmt.Println("Titulo: ", libro.Titulo)
	fmt.Println("Autor: ", libro.Autor)
}
