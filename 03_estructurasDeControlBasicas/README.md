---
id: 03-estructuras
title: Estructuras de control basicas
sidebar_label: Estructuras de control basicas
previous_page: 02-variables
next_page: 04-funciones
---
- [Estructuras de control básicas](#estructuras-de-control-básicas)
    - [Bucle FOR](#bucle-for)
    - [Bucle Range](#bucle-range)
    - [If](#if)
      - [Ejemplo básico de if](#ejemplo-básico-de-if)
      - [If con declaración de inicialización](#if-con-declaración-de-inicialización)
      - [If-else](#if-else)
    - [Switch](#switch)
      - [Ejemplo básico de switch](#ejemplo-básico-de-switch)
      - [Switch con múltiples valores en un caso](#switch-con-múltiples-valores-en-un-caso)
      - [Switch sin expresión](#switch-sin-expresión)
      - [Switch con declaración de inicialización](#switch-con-declaración-de-inicialización)
    - [Error Handling (agarrar los errores)](#error-handling-agarrar-los-errores)
      - [Declaración y manejo de errores en una función](#declaración-y-manejo-de-errores-en-una-función)
    - [Defer](#defer)
  - [Ejercicios](#ejercicios)
    - [Ejercicio 1: Sumar números del 1 al 100](#ejercicio-1-sumar-números-del-1-al-100)
    - [Ejercicio 2: Sumar elementos de un array](#ejercicio-2-sumar-elementos-de-un-array)
    - [Ejercicio 3: Determinar el día de la semana](#ejercicio-3-determinar-el-día-de-la-semana)
    - [Ejercicio 4: Manejo de errores al abrir un archivo](#ejercicio-4-manejo-de-errores-al-abrir-un-archivo)
    - [Ejercicio 5: Manejo de errores en una función de división](#ejercicio-5-manejo-de-errores-en-una-función-de-división)
    - [Ejercicio 6: Uso de defer para cerrar un archivo](#ejercicio-6-uso-de-defer-para-cerrar-un-archivo)

# Estructuras de control básicas

### Bucle FOR
* el for de Go no es muy diferente si miramos al de Java o Javascript 
* Tenemos 3 parámetros
  * el valor inicial
  * la condición
  * operación al valor inicial
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
    fmt.Println("valor de i: ", i)
    // valor de i:  0
    // valor de i:  1
    // ....
}
fmt.Println("total: ", sum)
// total:  45
```

### Bucle Range
* El `range` es utilizado para iterar `arrays`, `slices` y `maps` (tambien `canales`, pero todavía es pronto) en Go utilizando for:
* Tenemos 2 valores devueltos por range:
    * El `índice` del elemento.
    * El `valor` del elemento.
  
```go
numbers := []int{1, 2, 3, 4, 5}
sum := 0
// En el bucle for, el carácter _ se utiliza para ignorar el índice, ya que solo se necesita el valor del elemento para sumar los números.
for _, number := range numbers {
    sum += number
}
fmt.Println("La suma de los números es:", sum)

// con el indice:
for index, number := range numbers {
    fmt.Printf("Índice: %d, Valor: %d\n", index, number)
    sum += number
}
```

### If
El `if` en Go se utiliza para ejecutar bloques de código condicionalmente. La sintaxis básica es similar a otros lenguajes de programación, pero hay algunas diferencias notables:

* No es necesario el uso de paréntesis `()` alrededor de la condición.
* Las llaves `{}` son obligatorias incluso si el bloque `if` contiene una sola línea de código.
* Go soporta una declaración de inicialización antes de la condición en el `if`, la cual se ejecuta antes de evaluar la condición.

#### Ejemplo básico de if

```go
package main

import "fmt"

func main() {
    number := 10

    if number > 5 {
        fmt.Println("El número es mayor que 5")
    }
}
```

#### If con declaración de inicialización
```go
   if number := 10; number > 5 {
        fmt.Println("El número es mayor que 5")
    }
```

#### If-else
```go
   if number > 10 {
        fmt.Println("El número es mayor que 10")
    } else if number == 10 {
        fmt.Println("El número es 10")
    } else {
        fmt.Println("El número es menor que 10")
    }
```

### Switch
El `switch` en Go se utiliza para seleccionar entre múltiples bloques de código a ejecutar. Es más flexible y conciso que encadenar múltiples sentencias `if-else`.

* No es necesario el uso de `break` para salir de un caso, ya que Go lo hace automáticamente.
* Se pueden agrupar múltiples valores en un solo caso.
* El `switch` puede evaluar cualquier tipo de expresión, no solo valores enteros.
* Go soporta una declaración de inicialización antes de la expresión en el `switch`.

#### Ejemplo básico de switch

```go
package main

import "fmt"

func main() {
    day := "lunes"

    switch day {
    case "lunes":
        fmt.Println("Es lunes")
    case "martes":
        fmt.Println("Es martes")
    default:
        fmt.Println("No es ni lunes ni martes, tal vez miércoles, o jueves, quien sabe...")
    }
}
```

#### Switch con múltiples valores en un caso
```go
switch day {
  case "sábado", "domingo":
      fmt.Println("Es fin de semana")
  default:
      fmt.Println("Es un día de semana")
}
```

#### Switch sin expresión
```go
switch {
    case number < 0:
        fmt.Println("El número es negativo")
    case number == 0:
        fmt.Println("El número es cero")
}
```

#### Switch con declaración de inicialización
```go
switch number := 10; {
    case number < 0:
        fmt.Println("El número es negativo")
    case number == 0:
        fmt.Println("El número es cero")
}
```

### Error Handling (agarrar los errores)
* El manejo de errores en Go es explícito y se basa en la comprobación de valores de retorno de las funciones. En Go, las funciones que pueden fallar generalmente devuelven un valor y un error. Si el error no es `nil`, significa que ha ocurrido un error y debe ser manejado,
ej:

```Go
func main() {
    file, err := os.Open("archivo.txt")
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    fmt.Println("Archivo abierto con éxito")
}
```

#### Declaración y manejo de errores en una función
* Supongamos que estamos escribiendo una función que divide dos números y queremos manejar el caso donde el divisor es cero, lo cual generaría un error.

```Go
// divide dos números y devuelve el resultado y un error si el divisor es cero
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("el divisor no puede ser cero")
    }
    return a / b, nil
}

//main...
result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Resultado:", result)
    }
```

### Defer
* La palabra clave defer en Go se utiliza para posponer la ejecución de una función hasta que la función que la contiene haya terminado. Esto es muy útil para asegurar que ciertos recursos se liberen o ciertas acciones se realicen, sin importar cómo termine la función (ya sea de forma normal o debido a un error).

```Go
func main() {
    file, err := os.Open("archivo.txt")
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    // Asegura que el archivo se cierre al final de main, sin importar lo que ocurra.
    defer file.Close()

    fmt.Println("Archivo abierto con éxito")
}
```

## Ejercicios
- recuerda visitar la pagina [Golang playground](https://go.dev/play/) para testear y hacer el codigo.
### Ejercicio 1: Sumar números del 1 al 100
Escribe un programa que use un bucle `for` para sumar todos los números del 1 al 100 e imprima el resultado.
   - Si tienes problemas revisa el [ejercicio 1](./ejercicios/01.go)

### Ejercicio 2: Sumar elementos de un array
Declara un array de enteros nums con los valores {1, 2, 3, 4, 5}.
Usa un bucle range para sumar todos los elementos del array e imprime el resultado.
   - Si tienes problemas revisa el [ejercicio 2](./ejercicios/02.go)


### Ejercicio 3: Determinar el día de la semana
Declara una variable day de tipo string y asígnale un valor (por ejemplo, "lunes").
Usa una estructura switch para imprimir un mensaje dependiendo del valor de day.
   - Si tienes problemas revisa el [ejercicio 3](./ejercicios/03.go)

### Ejercicio 4: Manejo de errores al abrir un archivo
- Usa la función os.Open para intentar abrir un archivo llamado "archivo.txt".
- Maneja el error si el archivo no se puede abrir e imprime un mensaje de error.
- Si tienes problemas revisa el [ejercicio 4](./ejercicios/04.go)

### Ejercicio 5: Manejo de errores en una función de división
- Escribe una función divide que tome dos números a y b de tipo float64 y devuelva el resultado de la división y un error si el divisor es cero.
- Usa la función divide en el main para dividir 10 entre 0 y maneja el error apropiadamente.
- Si tienes problemas revisa el [ejercicio 5](./ejercicios/05.go)

### Ejercicio 6: Uso de defer para cerrar un archivo
- Usa la función os.Open para abrir un archivo llamado "archivo.txt".
- Utiliza defer para asegurarte de que el archivo se cierre al final del main, sin importar si ocurre un error.
- Si tienes problemas revisa el [ejercicio 6](./ejercicios/06.go)


<!-- [<< Anterior: Conceptos Básicos / variables ](../02_basico/README.md)
|
[Siguiente: Funciones y Paquetes >> ](../04_funcionesYPaquetes/README.md) -->