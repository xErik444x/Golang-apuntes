# Empecemos con lo bueno, se vienen los JSONs!

## ¿Qué es un JSON? te estarás preguntando... o no (?)
- Básicamente es un formato que se utiliza para intercambiar información de forma rápida y sencilla en la web.
- Se pueden usar para describir datos de forma legible y comprensible.
- Estos datos son llamados **objetos** o **estructuras**.

## ¿Como se utiliza este formato en Go?
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
	json, err := json.Marshal(p)
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
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Nombre)
}
```