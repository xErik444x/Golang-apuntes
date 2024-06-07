- [Funciones y paquetes](#funciones-y-paquetes)
  - [Funciones](#funciones)
    - [Ejemplo de una función simple](#ejemplo-de-una-función-simple)
    - [Ejemplo devolviendo dos valores](#ejemplo-devolviendo-dos-valores)
    - [Declarar funciones fuera del archivo principal](#declarar-funciones-fuera-del-archivo-principal)
  - [Paquetes](#paquetes)
    - [Creación de paquete custom](#creación-de-paquete-custom)
  - [Ejercicios!](#ejercicios)
      - [Ejercicio 1: Sumar dos números](#ejercicio-1-sumar-dos-números)
      - [Ejercicio 2: Devolver el doble y el triple de un número](#ejercicio-2-devolver-el-doble-y-el-triple-de-un-número)
      - [Ejercicio 3: Crear una función en otro archivo](#ejercicio-3-crear-una-función-en-otro-archivo)
      - [Ejercicio 4: Utiliza el paquete Math](#ejercicio-4-utiliza-el-paquete-math)
      - [Ejercicio 5: Crear un paquete personalizado](#ejercicio-5-crear-un-paquete-personalizado)

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
- `archivo matematicas.go`
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

## Paquetes
- Los paquetes son un conjunto de funciones que se pueden compartir entre archivos.
- Por ejemplo, `math` y `fmt` son paquetes.
- `fmt` es un paquete de funciones de escritura.
- `math` es un paquete de funciones matemáticas.
ej:
```go
package main
import "fmt"
import "math"
func main() {
    fmt.Println(math.Pi)
}
```
- Podemos obtener estos paquetes creándolos por nuestra cuenta o buscándolos en [pkg.go.dev](https://pkg.go.dev/)
- Primero debemos importarlo al proyecto con `go get rsc.io/quote`
- E importarlo en el main como hicimos con `fmt`, ej import `"rsc.io/quote"`

### Creación de paquete custom
- creamos una carpeta para guardar nuestro paquete custom dentro de la raiz del proyecto.
- creamos un archivo llamado `operaciones.go`
- dentro metemos la función `Sumar` y `Restar` siguientes:
```go
package operaciones

func Sumar(a int, b int) int {
    return a + b
}
func Restar(a int, b int) int {
    return a - b
}
```
- Si te das cuenta colocamos el nombre del paquete en la parte superior, ej `package operaciones`, y luego solo basta con importar el paquete colocando el identificador del proyecto (dentro del go.mod) y el nombre del paquete, ej:
```go
import "fmt"
import "idProyecto/descripcion/operaciones"
func main() {
    fmt.Println(operaciones.Sumar(2, 3))
    fmt.Println(operaciones.Restar(2, 3))
}
```
- Si querés indagar mas acerca de las dependencias / paquetes, puedes visitar [go.dev](https://go.dev/doc/modules/managing-dependencies)


## Ejercicios!
#### Ejercicio 1: Sumar dos números
1. Escribe una función llamada `sumar` que tome dos enteros como parámetros y devuelva su suma.
2. Llama a esta función desde `main` e imprime el resultado.
- si tienes problemas con este ejercicio puedes revisar el [ejercicio 1 main.go](ejercicios/01/main.go)

#### Ejercicio 2: Devolver el doble y el triple de un número
1. Escribe una función llamada `dobleTriple` que tome un `entero` como parámetro y `devuelva dos valores`: el `doble` y el `triple` del número.
2. Llama a esta función desde main e imprime los resultados.
- si tienes problemas con este ejercicio puedes revisar el [ejercicio 2 main.go](ejercicios/02/main.go)

#### Ejercicio 3: Crear una función en otro archivo
1. Crea una función `restar` que tome `dos enteros` y devuelva su diferencia en un archivo `matematicas.go`.
2. Llama a esta función desde main e imprime los resultados.
- si tienes problemas con este ejercicio puedes revisar el [ejercicio 3 main.go](ejercicios/03/main.go)
 [ejercicio 3 matematicas.go](ejercicios/03/matematicas.go)

#### Ejercicio 4: Utiliza el paquete Math
1. Usa el paquete `math` para calcular la raíz cuadrada de un número en el archivo `main.go`.
2. Imprime el resultado.
- si tienes problemas con este ejercicio puedes revisar el [ejercicio 4 main.go](ejercicios/04/main.go)

#### Ejercicio 5: Crear un paquete personalizado
1. Crea un paquete llamado `operaciones` que contenga una función `Multiplicar` que tome dos `enteros` y `devuelva su multiplicación o producto`.
2. Usa este paquete en `main.go` para multiplicar dos números e imprimir el resultado.
- si tienes problemas con este ejercicio puedes revisar el [ejercicio 5 main.go](ejercicios/05/main.go) y [ejercicio 5 operaciones.go](ejercicios/05/operaciones/operaciones.go)


[<< Anterior: Estructuras de control básicas](../03_estructurasDeControlBasicas/README.md)
|
[Siguiente: Type Casting \& Maps \& Slices \& Structs \& Pointers >> ](../05_typeCastingMapsSlicesStructsPointers/README.md)