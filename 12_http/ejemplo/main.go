package main

import (
	"fmt"
	"net/http"
)

// Manejador para la ruta "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, Mundo!")
}

func main() {
	http.HandleFunc("/", homeHandler) // Asocia la ruta "/" con el manejador homeHandler
	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil) // Inicia el servidor en el puerto 8080
}
