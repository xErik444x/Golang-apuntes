# Archivos, ficheros, y directorios
- Para escribir en un archivo, podés usar el paquete `os` para crear y abrir archivos y el paquete `bufio` o `io/ioutil` para escribir datos en ellos.

## Crear un archivo
- Para crear un archivo, puedes usar la función `os.Create`:
```go
f, err := os.Create("archivo.txt")
if err != nil {
    fmt.Println("Error al crear el archivo:", err)
    return
}
```