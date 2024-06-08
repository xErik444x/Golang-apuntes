# Qué es esta "guía"?
- Antes de nada, parece una guía extensa, lo es, pero no lo es del todo, solo cubro las características básicas y la sintaxis para entender Golang!
- Esta guía está hecha para aprender desde 0 las cosas más relevantes a la hora de aprender Golang, en especial si nunca lo viste o si sos principiante.
¿Por qué? Porque quería aprender Go y, al igual que hice con [Python](https://github.com/xErik444x/apuntesPython) , me pareció una buena idea crear una mini guía con lo esencial que, con suerte, pueda servirle a alguien en el futuro.

  ## Índice
- [Introducción a Golang](01_introduccion/README.md)
  - [Índice](01_introduccion/README.md#índice)
  - [Qué pingo es Golang?](01_introduccion/README.md#qué-pingo-es-golang)
  - [Para qué sirve??](01_introduccion/README.md#para-qué-sirve)
  - [Ejemplo de Código en Go](01_introduccion/README.md#ejemplo-de-código-en-go)
  - [Todo muy lindo, pero como lo ejecuto y uso?](01_introduccion/README.md#todo-muy-lindo-pero-como-lo-ejecuto-y-uso)
  - [Ejercicios!](01_introduccion/README.md#ejercicios)
- [Conceptos Básicos / variables](02_basico/README.md)
    - [Variables y declaraciones](02_basico/README.md#variables-y-declaraciones)
    - [tipos de datos variables](02_basico/README.md#tipos-de-datos-variables)
    - [Ejercicios!](02_basico/README.md##ejercicios)
- [estructuras de control básicas](03_estructurasDeControlBasicas/README.md#estructuras-de-control-básicas)
    - [Bucle FOR](03_estructurasDeControlBasicas/README.md#bucle-for)
    - [Bucle Range](03_estructurasDeControlBasicas/README.md#bucle-range)
    - [If](03_estructurasDeControlBasicas/README.md#if)
      - [Ejemplo básico de if](03_estructurasDeControlBasicas/README.md#ejemplo-básico-de-if)
      - [If con declaración de inicialización](03_estructurasDeControlBasicas/README.md#if-con-declaración-de-inicialización)
      - [If-else](03_estructurasDeControlBasicas/README.md#if-else)
    - [Switch](03_estructurasDeControlBasicas/README.md#switch)
      - [Ejemplo básico de switch](03_estructurasDeControlBasicas/README.md#ejemplo-básico-de-switch)
      - [Switch con múltiples valores en un caso](03_estructurasDeControlBasicas/README.md#switch-con-múltiples-valores-en-un-caso)
      - [Switch sin expresión](03_estructurasDeControlBasicas/README.md#switch-sin-expresión)
      - [Switch con declaración de inicialización](#switch-con-declaración-de-inicialización)
    - [Error Handling (agarrar los errores)](03_estructurasDeControlBasicas/README.md#error-handling-agarrar-los-errores)
      - [Declaración y manejo de errores en una función](03_estructurasDeControlBasicas/README.md#declaración-y-manejo-de-errores-en-una-función)
    - [Defer](03_estructurasDeControlBasicas/README.md#defer)
  - [Ejercicios](#ejercicios)
    - [Ejercicio 1: Sumar números del 1 al 100](03_estructurasDeControlBasicas/README.md#ejercicio-1-sumar-números-del-1-al-100)
    - [Ejercicio 2: Sumar elementos de un array](03_estructurasDeControlBasicas/README.md#ejercicio-2-sumar-elementos-de-un-array)
    - [Ejercicio 3: Determinar el día de la semana](03_estructurasDeControlBasicas/README.md#ejercicio-3-determinar-el-día-de-la-semana)
    - [Ejercicio 4: Manejo de errores al abrir un archivo](03_estructurasDeControlBasicas/README.md#ejercicio-4-manejo-de-errores-al-abrir-un-archivo)
    - [Ejercicio 5: Manejo de errores en una función de división](03_estructurasDeControlBasicas/README.md#ejercicio-5-manejo-de-errores-en-una-función-de-división)
    - [Ejercicio 6: Uso de defer para cerrar un archivo](#ejercicio-6-uso-de-defer-para-cerrar-un-archivo)
- [Funciones y paquetes](04_funcionesYPaquetes/README.md#funciones-y-paquetes)
  - [Funciones](04_funcionesYPaquetes/README.md#funciones)
    - [Ejemplo de una función simple](04_funcionesYPaquetes/README.md#ejemplo-de-una-función-simple)
    - [Ejemplo devolviendo dos valores](04_funcionesYPaquetes/README.md#ejemplo-devolviendo-dos-valores)
    - [Declarar funciones fuera del archivo principal](04_funcionesYPaquetes/README.md#declarar-funciones-fuera-del-archivo-principal)
  - [Paquetes](04_funcionesYPaquetes/README.md#paquetes)  - [Paquetes](#paquetes)
    - [Creación de paquete custom](04_funcionesYPaquetes/README.md#creación-de-paquete-custom)
  - [Ejercicios!](04_funcionesYPaquetes/README.md#ejercicios)
      - [Ejercicio 1: Sumar dos números](04_funcionesYPaquetes/README.md#ejercicio-1-sumar-dos-números)
      - [Ejercicio 2: Devolver el doble y el triple de un número](04_funcionesYPaquetes/README.md#ejercicio-2-devolver-el-doble-y-el-triple-de-un-número)
      - [Ejercicio 3: Crear una función en otro archivo](04_funcionesYPaquetes/README.md#ejercicio-3-crear-una-función-en-otro-archivo)
      - [Ejercicio 4: Utiliza el paquete Math](04_funcionesYPaquetes/README.md#ejercicio-4-utiliza-el-paquete-math)
      - [Ejercicio 5: Crear un paquete personalizado](04_funcionesYPaquetes/README.md#ejercicio-5-crear-un-paquete-personalizado)
- [Type Casting \& Maps \& Slices \& Structs \& Pointers](05_typeCastingMapsSlicesStructsPointers/README.md#type-casting--maps--slices--structs--pointers)
  - [Type casting](05_typeCastingMapsSlicesStructsPointers/README.md#type-casting)
    - [Ejemplo de Type Casting:](05_typeCastingMapsSlicesStructsPointers/README.md#ejemplo-de-type-casting)
  - [Maps](05_typeCastingMapsSlicesStructsPointers/README.md#maps)
  - [Slices](05_typeCastingMapsSlicesStructsPointers/README.md#slices)
  - [Structs](05_typeCastingMapsSlicesStructsPointers/README.md#structs)
  - [make](05_typeCastingMapsSlicesStructsPointers/README.md#make)
    - [Slices](05_typeCastingMapsSlicesStructsPointers/README.md#slices-1)
    - [Maps](05_typeCastingMapsSlicesStructsPointers/README.md#maps-1)
  - [Pointers](05_typeCastingMapsSlicesStructsPointers/README.md#pointers)
    - [¿Qué es un Puntero?](05_typeCastingMapsSlicesStructsPointers/README.md#qué-es-un-puntero)
    - [¿Cómo se crea un puntero?](05_typeCastingMapsSlicesStructsPointers/README.md#como-se-crea-un-puntero)
    - [Inicialización de Punteros](05_typeCastingMapsSlicesStructsPointers/README.md#inicialización-de-punteros)
    - [Punteros en Funciones](05_typeCastingMapsSlicesStructsPointers/README.md#punteros-en-funciones)
    - [Punteros a Estructuras (Structs)](05_typeCastingMapsSlicesStructsPointers/README.md#punteros-a-estructuras-structs)
    - [Nil y Punteros](05_typeCastingMapsSlicesStructsPointers/README.md#nil-y-punteros)
  - [Ejercicios!](05_typeCastingMapsSlicesStructsPointers/README.md#ejercicios)
    - [Ejercicio 1: Type Casting](05_typeCastingMapsSlicesStructsPointers/README.md#ejercicio-1-type-casting)
    - [Ejercicio 2: estudiantes Maps](05_typeCastingMapsSlicesStructsPointers/README.md#ejercicio-2-estudiantes-maps)
    - [Ejercicio 3: Slices](05_typeCastingMapsSlicesStructsPointers/README.md#ejercicio-3-slices)
    - [Ejercicio 4: Libros! (Structs)](05_typeCastingMapsSlicesStructsPointers/README.md#ejercicio-4-libros-structs)
    - [Ejercicio 5: Pointers](05_typeCastingMapsSlicesStructsPointers/README.md#ejercicio-5-pointers)
- [Introducción a la Creación de CLI y Compilaciones](06_CLI_basics/README.md#introducción-a-la-creación-de-cli-y-compilaciones)
  - [Uso de la Biblioteca flag](06_CLI_basics/README.md#uso-de-la-biblioteca-flag)
    - [Quiero generar un ejecutable del programa, cómo hago?](06_CLI_basics/README.md#quiero-generar-un-ejecutable-del-programa-como-hago)
    - [compilación cruzada:](06_CLI_basics/README.md#compilación-cruzada)
      - [Listar las plataformas disponibles:](06_CLI_basics/README.md#listar-las-plataformas-disponibles)
      - [Compilar:](06_CLI_basics/README.md#compilar)
      - [compilar en Linux:](06_CLI_basics/README.md#compilar-en-linux)
      - [compilar en Windows:](06_CLI_basics/README.md#compilar-en-windows)
      - [compilar en Macos:](06_CLI_basics/README.md#compilar-en-macos)
      - [Ejercicios!](06_CLI_basics/README.md#ejercicios)
        - [Crear un CLI que multiplique dos números](06_CLI_basics/README.md#crear-un-cli-que-multiplique-dos-números)
- [Archivos, ficheros, y directorios](07_archivos/README.md#archivos-ficheros-y-directorios)
    - [Crear un archivo](07_archivos/README.md#crear-un-archivo)
    - [Escribir en un archivo](07_archivos/README.md#escribir-en-un-archivo)
    - [Leer un archivo](07_archivos/README.md#leer-un-archivo)
    - [Borrar un archivo](07_archivos/README.md#borrar-un-archivo)
    - [Métodos destacables del paquete `os` para archivos y carpetas](07_archivos/README.md#métodos-destacables-del-paquete-os-para-archivos-y-carpetas)
  - [Ejercicios!](07_archivos/README.md#ejercicios)
  - [Mazmorra de los Archivos](07_archivos/README.md#mazmorra-de-los-archivos)
    - [Nivel 1: La Sala de Creación](07_archivos/README.md#nivel-1-la-sala-de-creación)
    - [Nivel 2: La Sala de la Escritura](07_archivos/README.md#nivel-2-la-sala-de-la-escritura)
    - [Nivel 3: La Sala de la Lectura](07_archivos/README.md#nivel-3-la-sala-de-la-lectura)
    - [Nivel 4: La Sala del Borrado](07_archivos/README.md#nivel-4-la-sala-del-borrado)
    - [Nivel 5: La Sala de los Métodos Especiales](07_archivos/README.md#nivel-5-la-sala-de-los-métodos-especiales)