package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	aventura, err := os.Create("aventura.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo :C ", err)
		return
	}
	defer aventura.Close()
	time.Sleep(1 * time.Second)
	_, err = aventura.WriteString("¡Bienvenido a la Mazmorra de los Archivos!")
	if err != nil {
		fmt.Println("Error al escribir en el archivo :c ", err)
		return
	}
	time.Sleep(1 * time.Second)
	file, err := os.ReadFile("aventura.txt")
	if err != nil {
		fmt.Println("Error al leer en el archivo :c ", err)
		return
	}
	fmt.Println(string(file))
	time.Sleep(1 * time.Second)
	err = os.Mkdir("mazmorra", 0755)
	if err != nil {
		fmt.Println("Error al crear carpeta mazmorra :c ", err)
		return
	}
	time.Sleep(2 * time.Second)
	err = os.Rename("mazmorra", "mazmorra_final")
	if err != nil {
		fmt.Println("Error al renombrar carpeta mazmorra :c ", err)
		return
	}
	time.Sleep(2 * time.Second)
	err = os.Remove("mazmorra_final")
	if err != nil {
		fmt.Println("Error al borrar carpeta mazmorra :c ", err)
		return
	}
}
