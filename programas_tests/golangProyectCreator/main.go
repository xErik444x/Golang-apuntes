package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: gop <nombre del proyecto>")
		return
	}
	projectName := os.Args[1]
	nameForFolder := strings.ReplaceAll(projectName, " ", "_")
	nameForFolder = strings.ReplaceAll(nameForFolder, "/", "_")
	fmt.Printf("Creando proyecto %s...\n", projectName)

	errFolder := os.Mkdir("./"+nameForFolder, 0755)
	if errFolder != nil {
		fmt.Println(errFolder)
		return
	}

	fmt.Printf("Se creó la carpeta ./%s con éxito. Accediendo y ejecutando Go mod init %s", nameForFolder, projectName)

	errGoMod := os.WriteFile(fmt.Sprintf("./%s/go.mod", nameForFolder), []byte(fmt.Sprintf("module %s\n", projectName)), 0644)
	if errGoMod != nil {
		fmt.Println(errGoMod)
		return
	}
	errMainGo := os.WriteFile(fmt.Sprintf("./%s/main.go", nameForFolder), []byte("package main\n\nfunc main() {\n\n}\n"), 0644)
	if errMainGo != nil {
		fmt.Println(errMainGo)
		return
	}

	fmt.Println("Proyecto creado con éxito.")
	fmt.Println("Iniciando vs code...")

	errCode := os.Chdir(fmt.Sprintf("./%s", nameForFolder))
	if errCode != nil {
		fmt.Println(errCode)
		return
	}
	exec.Command("code", ".").Run()
}
