---
id: 09-json
title: JSON!
sidebar_label: JSON!
previous_page: 08-ejercicios
next_page: 10-interfaces
---
- [Empecemos con lo bueno, se vienen los JSONs!](#empecemos-con-lo-bueno-se-vienen-los-jsons)
	- [¿Qué es un JSON? te estarás preguntando... o no (?)](#qué-es-un-json-te-estarás-preguntando-o-no-)
	- [¿Cómo se utiliza este formato en Go?](#cómo-se-utiliza-este-formato-en-go)
		- [Manejo de Estructuras Anidadas](#manejo-de-estructuras-anidadas)
		- [Etiquetas!](#etiquetas)
		- [Recomendación NO SALTAR!](#recomendación-no-saltar)

# Empecemos con lo bueno, se vienen los JSONs!

## ¿Qué es un JSON? te estarás preguntando... o no (?)
- Básicamente es un formato que se utiliza para intercambiar información de forma rápida y sencilla en la web.
- Se pueden usar para describir datos de forma legible y comprensible.
- Estos datos son llamados **objetos** o **estructuras**.

## ¿Cómo se utiliza este formato en Go?
- Lo ideal es tener a mano la url del paquete: [encoding/json package](https://pkg.go.dev/encoding/json#pkg-overview)
- Para empezar necesitamos utilizar la función `Marshal` y `Unmarshal` del paquete `encoding/json`.
- La función `Marshal` convierte un objeto de Go en un objeto JSON.
- La función `Unmarshal` convierte un objeto JSON en un objeto de Go.
- ej de objeto de Go a JSON: [Url para testear online](https://goplay.tools/snippet/yBDcmNlTo4x)
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string
}

func main() {
	p := Persona{Nombre: "Juan"}
	json, err := json.Marshal(p) // Marshal convierte un objeto de Go en un objeto JSON
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
}
```
- ej de JSON a objeto de Go: [Url para testear online](https://goplay.tools/snippet/KffxhrFVmsA)
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string
}

func main() {
	jsonStr := `{"Nombre": "Juan"}`
	p := Persona{}
	err := json.Unmarshal([]byte(jsonStr), &p) // Unmarshal convierte un objeto JSON en un objeto de Go
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Nombre)
}
```

### Manejo de Estructuras Anidadas
- En caso de que un objeto JSON contenga una estructura anidada, la función `Unmarshal` puede ser utilizada para convertir el objeto JSON en un objeto de Go.
- La idea es hacer la estructura completa del JSON para mejorar el manejo y evitar problemas de codificación y decodificación.
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Direccion struct {
	Calle  string
	Ciudad string
}

type Persona struct {
	Nombre    string
	Direccion Direccion
}

func main() {
	p := Persona{
		Nombre: "Juan",
		Direccion: Direccion{
			Calle:  "Calle Falsa 123",
			Ciudad: "Springfield",
		},
	}
	json, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json)) 
	// {"Nombre":"Juan","Direccion":{"Calle":"Calle Falsa 123","Ciudad":"Springfield"}}
}

```
### Etiquetas!
- Al definir una estructura, podes definir etiquetas para los atributos.
- Estas etiquetas pueden usarse para cambiar el nombre de un atributo en el json o para hacer alguna combinación / comprobación.
- ej de etiquetas:
```go
type Persona struct {
	Nombre string `json:"name"`
	// cuando guardemos persona en el json, en vez de guardarse como {"Nombre":"Juan"}, guardaremos {"name": "Juan"}
}
```
```go
type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad,omitempty"`
	//en este caso, si la edad es vacia, no la guardamos
}
```
- Y ahora un ejemplo con todas las etiquetas:
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Direccion struct {
	Calle  string `json:"calle"`
	Ciudad string `json:"ciudad,omitempty"`
}

type Persona struct {
	Nombre     string    `json:"nombre"`
	Edad       int       `json:"edad,omitempty"`
	Contraseña string    `json:"-"` // Campo omitido siempre
	Direccion  Direccion `json:"direccion"`
}

func main() {
	p := Persona{
		Nombre: "Juan",
		Edad:   0, // Este campo será omitido en el JSON debido a la etiqueta `omitempty`
		Direccion: Direccion{
			Calle:  "Calle Falsa 123",
			Ciudad: "",
		},
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
	// Resultado: {"nombre":"Juan","direccion":{"calle":"Calle Falsa 123"}}
}
```

### Recomendación NO SALTAR!
- Para obtener estos structs y no hacerlos a mano, podés usar el sitio: [JSON to Go](https://transform.tools/json-to-go), pegando el json te devuelve ya las estructuras ;)
