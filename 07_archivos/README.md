---
id: 07-archivos
title: Archivos!
sidebar_label: Archivos!
---
- [Archivos, ficheros, y directorios](#archivos-ficheros-y-directorios)
		- [Crear un archivo](#crear-un-archivo)
		- [Escribir en un archivo](#escribir-en-un-archivo)
		- [Leer un archivo](#leer-un-archivo)
		- [Borrar un archivo](#borrar-un-archivo)
		- [Métodos destacables del paquete `os` para archivos y carpetas](#métodos-destacables-del-paquete-os-para-archivos-y-carpetas)
	- [Ejercicios!](#ejercicios)
	- [Mazmorra de los Archivos](#mazmorra-de-los-archivos)
		- [Nivel 1: La Sala de Creación](#nivel-1-la-sala-de-creación)
		- [Nivel 2: La Sala de la Escritura](#nivel-2-la-sala-de-la-escritura)
		- [Nivel 3: La Sala de la Lectura](#nivel-3-la-sala-de-la-lectura)
		- [Nivel 4: La Sala del Borrado](#nivel-4-la-sala-del-borrado)
		- [Nivel 5: La Sala de los Métodos Especiales](#nivel-5-la-sala-de-los-métodos-especiales)

# Archivos, ficheros, y directorios
- Para escribir en un archivo, podés usar el paquete `os` para crear y abrir archivos y el paquete `bufio` o `io/ioutil` para escribir datos en ellos.

### Crear un archivo
- Para crear un archivo, puedes usar la función `os.Create`, esta devuelve el archivo abierto y un error (si hay):
```go
	archivo, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error al crear el archivo :c ", err)
	}
	defer archivo.Close() // debe cerrarse al final de la ejecución para evitar errores / limpiar recursos.
```
### Escribir en un archivo
- Para escribir en un archivo, puedes usar el paquete `os` para crear y abrir archivos, una vez abierto llamas a `WriteString` con el texto que deseas escribir y listo ;).
```go
	archivo, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error al crear el archivo :c ", err)
	}
	defer archivo.Close()

	texto := "Archivos desbloqueados!"
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("error al escribir en el archivo >:C ", err)
	}
}
```
### Leer un archivo
- Para leer un archivo, puedes usar el paquete `os` con la función `ReadFile` y lesto papurri, ahí tenes tu archivo todo lindo, todo bonito.
```go
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error al leer el archivo D:", err)
		return
	}
	fmt.Println(string(content))
}
```
### Borrar un archivo
- Para borrar un archivo, puedes usar el paquete `os` con la función `Remove`. Así, seco, sin nada más que añadir >:C
```go
    err = os.Remove("test.txt")
	if err != nil {
		fmt.Println("error al borrar el archivo D:", err)
	}
```
### Métodos destacables del paquete `os` para archivos y carpetas
- `os.Mkdir` para crear carpetas
- `os.MkdirAll` para crear carpetas con subcarpetas
- `os.Rename` para renombrar archivos
- `os.RemoveAll` para borrar carpetas
- `os.Remove` para borrar archivos
- `os.Stat` para saber si un archivo existe
- `os.Chmod` para cambiar permisos de archivos
- `os.Chdir` para cambiar de directorio
- `os.Chtimes` para cambiar la fecha de acceso y modificación de un archivo
- `os.Symlink` para crear un enlace simbólico
- `os.Readlink` para leer el enlace simbólico
- `os.TempDir` para crear un directorio temporal
- `os.TempFile` para crear un archivo temporal
- `os.Lstat` para saber si un archivo es un enlace simbólico
- `os.Link` para crear un enlace simbólico de un archivo
- `os.SameFile` para saber si dos archivos son iguales
- `os.Open` para abrir un archivo

> Por las dudas loco, revisá el  [paquete `os`](https://pkg.go.dev/os)

## Ejercicios!
Vamos a cambiar un poco la dinámica, que te parece si hacemos una especie de mazmorra? esta vez no te voy a dar el código, disfrutá pensar y programar loquito.

## Mazmorra de los Archivos

### Nivel 1: La Sala de Creación
**Desafío: Crear un archivo "aventura.txt"**

### Nivel 2: La Sala de la Escritura
**Desafío: Escribir "¡Bienvenido a la Mazmorra de los Archivos!" en "aventura.txt"**

### Nivel 3: La Sala de la Lectura
**Desafío: Leer el contenido de "aventura.txt" y mostrarlo en la consola**

### Nivel 4: La Sala del Borrado
**Desafío: Borrar "aventura.txt"**

### Nivel 5: La Sala de los Métodos Especiales
**Desafío: Crear un directorio "mazmorra", renombrarlo a "mazmorra_final" y luego borrarlo**


[<< Anterior: CLI Basics](../06_CLI_basics/README.md)

