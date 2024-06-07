- [Type Casting \& Maps \& Slices \& Structs \& Pointers](#type-casting--maps--slices--structs--pointers)
  - [Type casting](#type-casting)
    - [Ejemplo de Type Casting:](#ejemplo-de-type-casting)
  - [Maps](#maps)
  - [Slices](#slices)
  - [Structs](#structs)
  - [make](#make)
    - [Slices](#slices-1)
    - [Maps](#maps-1)
  - [Pointers](#pointers)
    - [¿Qué es un Puntero?](#qué-es-un-puntero)
    - [¿Como se crea un puntero?](#como-se-crea-un-puntero)
    - [Inicialización de Punteros](#inicialización-de-punteros)
    - [Punteros en Funciones](#punteros-en-funciones)
    - [Punteros a Estructuras (Structs)](#punteros-a-estructuras-structs)
    - [Nil y Punteros](#nil-y-punteros)

# Type Casting & Maps & Slices & Structs & Pointers

## Type casting
- En Go, la conversión de tipos (type casting) se realiza explícitamente usando la sintaxis `T(v)`, donde `T` es el tipo al que se quiere convertir y `v` es el valor a convertir.

### Ejemplo de Type Casting:
```go
var x int = 42
var y float64 = float64(x) // convertimos x a float64
fmt.Println(y) // 42.0
```
## Maps
- Los maps son colecciones desordenadas de pares clave/valor. Las claves deben ser de un tipo que sea comparable, y los valores pueden ser de cualquier tipo.
- los podemos ver como si fueran un objeto de javascript.
```go
var m map[string]int = map[string]int{"valor_uno": 1, "valor_dos": 2}
fmt.Println(m)
```

## Slices
- Los slices son más flexibles que los arrays y pueden cambiar de tamaño dinámicamente. Son construidos sobre arrays.
```go
var slice []int = []int{1, 2, 3}
fmt.Println(slice)
```

## Structs
- Las structs son tipos de datos que agrupan campos. Cada campo tiene un nombre y un tipo.
- Los structs funcionan como contenedores de datos estructurados, similares a clases en otros lenguajes, pero sin métodos o herencia.
```go
type Persona struct {
    Nombre string
    Edad   int
}

func main() {
    var p Persona = Persona{"Juan", 30}
    fmt.Println(p)        // Imprime: {Juan 30}
    fmt.Println(p.Edad)   // Imprime: 30
}
```

## make

La función `make` en Go se utiliza para inicializar `slices`, `maps` y `canales`. Es útil cuando necesitas crear estos tipos de datos con una longitud y capacidad específicas.

### Slices
Para crear un slice con una longitud y capacidad específica.

```go
func main() {
    slice := make([]int, 5, 10)
    fmt.Println(slice)       // Imprime: [0 0 0 0 0]
    fmt.Println(len(slice))  // Imprime: 5 (longitud)
    fmt.Println(cap(slice))  // Imprime: 10 (capacidad)
}
```

### Maps
```go
func main() {
    m := make(map[string]int)
    m["key"] = 42
    fmt.Println(m["key"])  // Imprime: 42
}
```

## Pointers
### ¿Qué es un Puntero?
- Un puntero es una variable que almacena la dirección de memoria de otra variable.

### ¿Como se crea un puntero?
- Se utiliza el operador `*` para declarar un puntero.
ej:
```go
var p *int
```
### Inicialización de Punteros
- Se inicializa el puntero con la variable que deseas convertir a puntero.
- Ejemplo:
```go
var x int = 42
var p *int = &x
fmt.Println(*p) // Imprime: 42
```
- Se usa el operador * para acceder o modificar el valor apuntado.
```go
var x int = 42
var p *int = &x
fmt.Println(*p) // Imprime: 42
*p = 20          // Cambia el valor de x a 20
fmt.Println(x)   // Imprime 20
```
### Punteros en Funciones
- Pasar punteros permite modificar / obtener el valor original de las variables.
```go
func modificarValor(p *int) {
    *p = 30
}

func main() {
    var x int = 10
    modificarValor(&x)
    fmt.Println(x)  // Imprime 30
}
```

### Punteros a Estructuras (Structs)
- Los punteros pueden apuntar a estructuras para evitar copias costosas.
```go
type Persona struct {
    nombre string
    edad   int
}

func main() {
    p := Persona{"Juan", 25}
    var ptr *Persona = &p

    fmt.Println(ptr.nombre)  // Acceso al campo nombre mediante el puntero
    ptr.edad = 26            // Modificación del campo edad mediante el puntero
    fmt.Println(p.edad)      // Imprime 26
}
```
### Nil y Punteros
- Los punteros pueden ser `nil`, que significa que no apunta a ninguna variable.
```go
var p *int
if p != nil {
    fmt.Println(*p)
} else {
    fmt.Println("El puntero es nil")
}
```