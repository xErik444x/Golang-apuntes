package utils

import (
	"bufio"
	"fmt"
	"log"
	"mysql/01/handlers"
	"mysql/01/models"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/fatih/color"
)

func PrintTitle() {
	log.Println(color.BlueString(`

    _____ ____    __       ________    ___________   _____________
   / ___// __ \  / /      / ____/ /   /  _/ ____/ | / /_  __/ ___/
   \__ \/ / / / / /      / /   / /    / // __/ /  |/ / / /  \__ \ 
  ___/ / /_/ / / /___   / /___/ /____/ // /___/ /|  / / /  ___/ / 
 /____/\___\_\/_____/   \____/_____/___/_____/_/ |_/ /_/  /____/  
   `))
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Startup() {
	var ID int
	var name, email, tel string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(color.CyanString("Presione una tecla para continuar..."))
		if !scanner.Scan() {
			break
		}
		fmt.Println(color.CyanString("\n"))
		PrintTitle()
		fmt.Println(color.CyanString("Seleccione una opción:"))
		fmt.Println(color.CyanString("1. Listar Clientes"))
		fmt.Println(color.CyanString("2. Listar Cliente por id"))
		fmt.Println(color.CyanString("3. Agregar Clientes"))
		fmt.Println(color.CyanString("4. Editar Clientes"))
		fmt.Println(color.CyanString("5. Eliminar Clientes"))
		fmt.Println(color.CyanString("6. Salir"))
		if !scanner.Scan() {
			break
		}
		switch scanner.Text() {
		case "1":
			handlers.GetClients()
		case "2":
			fmt.Print(color.BlueString("Ingrese el ID del cliente: "))
			if scanner.Scan() {
				ID, _ = strconv.Atoi(scanner.Text())
			}
			handlers.GetClientById(ID)
		case "3":
			fmt.Print(color.BlueString("Ingrese el NOMBRE del cliente: "))
			if scanner.Scan() {
				name = scanner.Text()
			}
			fmt.Print(color.BlueString("Ingrese el EMAIL del cliente: "))
			if scanner.Scan() {
				email = scanner.Text()
			}
			fmt.Print(color.BlueString("Ingrese el TELEFONO del cliente: "))
			if scanner.Scan() {
				tel = scanner.Text()
			}
			cliente := models.Cliente{Id: ID, Nombre: name, Correo: email, Telefono: tel}
			handlers.InsertClient(cliente)
			// GetClientById(ID)
		case "4":
			fmt.Print(color.BlueString("Ingrese el ID del cliente: "))
			if scanner.Scan() {
				ID, _ = strconv.Atoi(scanner.Text())
			}
			fmt.Print(color.BlueString("Ingrese el NOMBRE del cliente: "))
			if scanner.Scan() {
				name = scanner.Text()
			}
			fmt.Print(color.BlueString("Ingrese el EMAIL del cliente: "))
			if scanner.Scan() {
				email = scanner.Text()
			}
			fmt.Print(color.BlueString("Ingrese el TELEFONO del cliente: "))
			if scanner.Scan() {
				tel = scanner.Text()
			}
			cliente := models.Cliente{Id: ID, Nombre: name, Correo: email, Telefono: tel}
			handlers.EditClient(ID, cliente)
		case "5":
			fmt.Print(color.BlueString("Ingrese el ID del cliente: "))
			if scanner.Scan() {
				ID, _ = strconv.Atoi(scanner.Text())
			}
			handlers.DeleteClient(ID)
		case "6":
			return
		default:
			fmt.Println(color.RedString("Opcion no valida"))
		}
	}
}
