package main

import (
	"fmt"
	"os"
)

func main() {
    file, err := os.Open("archivo.txt")
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    defer file.Close()
    
    fmt.Println("Archivo abierto con éxito")
}