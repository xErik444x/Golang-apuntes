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
    - [¿Cómo se crea un puntero?](#cómo-se-crea-un-puntero)
    - [Inicialización de Punteros](#inicialización-de-punteros)
    - [Punteros en Funciones](#punteros-en-funciones)
    - [Punteros a Estructuras (Structs)](#punteros-a-estructuras-structs)
    - [Nil y Punteros](#nil-y-punteros)
  - [Ejercicios!](#ejercicios)
    - [Ejercicio 1: Type Casting](#ejercicio-1-type-casting)
    - [Ejercicio 2: estudiantes Maps](#ejercicio-2-estudiantes-maps)
    - [Ejercicio 3: Slices](#ejercicio-3-slices)
    - [Ejercicio 4: Libros! (Structs)](#ejercicio-4-libros-structs)
    - [Ejercicio 5: Pointers](#ejercicio-5-pointers)

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

La función `make` se utiliza para inicializar `slices`, `maps` y `canales`. Es útil cuando necesitas crear estos tipos de datos con una longitud y capacidad específicas.

### Slices
Para crear un slice con una longitud y capacidad específica.
- cual es la diferencia entre longitud y capacidad? 
- La `longitud` de un slice representa la cantidad de elementos que contiene actualmente, mientras que la `capacidad` indica el tamaño máximo que puede alcanzar el slice antes de que necesite asignar más memoria.
```go
func main() {
    slice := make([]int, 5, 10)
    fmt.Println(slice)       // Imprime: [0 0 0 0 0]
    fmt.Println(len(slice))  // Imprime: 5 (longitud)
    fmt.Println(cap(slice))  // Imprime: 10 (capacidad)
}
```
- para agregar elementos a un slice:
```go
func main() {
    slice := make([]int, 5, 10)
    slice = append(slice, 1, 2, 3)
    fmt.Println(slice)       // Imprime: [0 0 0 0 0 1 2 3]
    fmt.Println(len(slice))  // Imprime: 8 (longitud)
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

### ¿Cómo se crea un puntero?
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

## Ejercicios!
### Ejercicio 1: Type Casting
- Convierte un `float64` a un `int` y un `int` a un `float64`. Calcula la suma de ambos y muestra el resultado en ambos tipos.
```go
    var a float64 = 7.5
    var b int = 3
    // completa el código
    
    fmt.Printf("Suma como float64: %.2f\n", sumFloat)
    fmt.Printf("Suma como int: %d\n", sumInt)
```
- Si tienes dudas o problemas revisa el [ejercicio 1](./ejercicios/01/main.go)

### Ejercicio 2: estudiantes Maps
- Crea un `map` para almacenar `nombres de estudiantes y sus calificaciones`. Agrega al menos tres estudiantes y sus calificaciones, luego muestra todas las calificaciones con un bucle Range.
- Si tienes dudas o no recuerdas como usar el range:  [Bucle Range](../03_estructurasDeControlBasicas/README.md#bucle-range)
- Si tienes dudas o problemas revisa el [ejercicio 2](./ejercicios/02/main.go)
  
### Ejercicio 3: Slices
- Crea un `slice` de `enteros` con `make` con al menos `cinco elementos`. `Agrega dos elementos` más al slice y `muestra la longitud y capacidad` del slice después de agregar los nuevos elementos.
- Si tienes dudas o no recuerdas como usar el Slice con Make:  [Make Slices](#slices-1)
- Si tienes dudas o problemas revisa el [ejercicio 3](./ejercicios/03/main.go)

### Ejercicio 4: Libros! (Structs)
- Define una `struct` llamada `Libro` con los campos `Titulo` y `Autor`. Crea una instancia de Libro, inicializa sus campos y muestra sus valores.
- Si tienes dudas o no recuerdas como usar Structs:  [Structs](#structs)
- Si tienes dudas o problemas revisa el [ejercicio 4](./ejercicios/04/main.go)

### Ejercicio 5: Pointers
- Crea una `función` que tome un `puntero` a un `int` y `duplique` el valor al que apunta. En main, define una variable entera, pasa su dirección a la función y muestra el valor antes y después de la llamada a la función.
- - Si tienes dudas o no recuerdas como usar Punteros:  [Punteros](#pointers)
- Si tienes dudas o problemas revisa el [ejercicio 5](./ejercicios/05/main.go)

[<< Anterior: Funciones y paquetes](../04_funcionesYPaquetes/README.md)
|
[Siguiente: CLI Basics >> ](../06_CLI_basics/README.md)