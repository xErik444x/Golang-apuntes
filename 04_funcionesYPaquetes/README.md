- [Funciones y paquetes](#funciones-y-paquetes)
  - [Funciones](#funciones)
    - [Ejemplo de una función simple](#ejemplo-de-una-función-simple)
    - [Ejemplo devolviendo dos valores](#ejemplo-devolviendo-dos-valores)
    - [Declarar funciones fuera del archivo principal](#declarar-funciones-fuera-del-archivo-principal)

# Funciones y paquetes
## Funciones
- las funciones de Go no difiere de las de Java o Javascript.
- la función main por ejemplo es la que se encarga de ejecutar el programa.
- tambien podemos declarar y usar funciones propias.
### Ejemplo de una función simple
```go
func main() {
    fmt.Println(sumar(2, 3))
} 

func sumar(a int, b int) int {
    return a + b
}
```
### Ejemplo devolviendo dos valores
```go
func main() {
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
} 
func vals() (int, int) {
    return 3, 7
}
```

### Declarar funciones fuera del archivo principal
- Las funciones también se pueden declarar fuera del archivo principal y ser utilizadas en el mismo o en otros archivos.
- `archivo math.go`
```go
package main

func sumar(a int, b int) int {
    return a + b
}
```
- `archivo main.go`
```go
package main

import "fmt"

func main() {
    fmt.Println(sumar(2, 3))
}
```
- Si te das cuenta, no hace falta importar la función, porque ya forma parte del paquete main.