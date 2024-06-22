package main

import (
	"fmt"
	"os"
)

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func writeFile(file *os.File, text string) error {
	_, err := file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func readFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func main() {
	fileName := "test.txt"

	archivo, err := createFile(fileName)
	if err != nil {
		fmt.Println("error al crear el archivo :c ", err)
		return
	}
	defer archivo.Close()

	texto := "Archivos desbloqueados!"
	err = writeFile(archivo, texto)
	if err != nil {
		fmt.Println("error al escribir en el archivo >:C ", err)
		return
	}

	content, err := readFile(fileName)
	if err != nil {
		fmt.Println("error al leer el archivo D:", err)
		return
	}
	fmt.Println(string(content))
}
