# Interfaces y Tipos!

## ¿Que es una interface?
- una interfaz es un tipo que define un conjunto de métodos, pero sin implementar ninguno de ellos. Cualquier tipo que implemente esos métodos se considera que implementa la interfaz.

### Conceptos rápidos a tener en cuenta:
- Interfaz:
  - Una interfaz define un conjunto de métodos sin implementar ninguno.
  - Cualquier tipo que implemente esos métodos satisface la interfaz.
- Método:
  - Un método es una función que está asociada a un tipo específico.

### Ejemplo básico de interfaz

**paso 1: Definir la interfaz**
> definimos una interfaz `Informable`, que define/contiene un método Informar que devuelve un string.
```go
type Informable interface {
    Informar() string
}
```
**paso 2:  Crear una Estructura (Tipo)**
> creamos una estructura `Persona` para implementar la interfaz `Informable`.
```go
type Persona struct {
    Nombre string
}
```
**Paso 3: Implementar el método `Informar` en la estructura `Persona`**
```go
func (p Persona) Informar() string {
    return "Hola, mi nombre es " + p.Nombre
}
```

**Paso 4: Crear un objeto de tipo `Persona` y llamar al método `Informar`**
```go
func main() {
    p := Persona{Nombre: "Juan"}
    fmt.Println(p.Informar())
}
```
