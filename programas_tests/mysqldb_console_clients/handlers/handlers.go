package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"mysql/01/connection"
	"mysql/01/models"

	"github.com/fatih/color"
)

// rutas de la api

func GetClients() {
	con := connection.Db

	sql := "SELECT id, nombre, correo, telefono from clientes order by id asc;"
	datos, _ := con.Query(sql)
	defer datos.Close()

	fmt.Println(color.CyanString("\n Listado de Clientes"))

	printTable(datos)
}

func GetClientById(id int) {
	fmt.Println(color.CyanString("\nListando Cliente por id: %v", id))
	con := connection.Db

	sql := "SELECT * FROM clientes WHERE id = ?"
	data, err := con.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	printTable(data)
}

func printTable(data *sql.Rows) {
	fmt.Printf(color.MagentaString("%-4s %-10s %-25s %-15s\n"), "ID", "Nombre", "Correo", "Telefono")
	for data.Next() {
		client := models.Cliente{}
		err := data.Scan(&client.Id, &client.Nombre, &client.Correo, &client.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(color.CyanString("%-4d %-10s %-25s %-15s\n"), client.Id, client.Nombre, client.Correo, client.Telefono)
	}
}

func InsertClient(cliente models.Cliente) {
	con := connection.Db
	sql := "INSERT INTO clientes values(null, ?, ?, ?);"
	_, err := con.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(color.CyanString("Se agrego el cliente"))
}

func EditClient(id int, cliente models.Cliente) {
	con := connection.Db
	sql := "UPDATE clientes SET nombre = ?, correo = ?, telefono = ? WHERE id = ?;"
	_, err := con.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(color.CyanString("Se edito el cliente"))
}

func DeleteClient(id int) {
	con := connection.Db
	sql := "DELETE FROM clientes WHERE id = ?;"
	_, err := con.Exec(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(color.CyanString("Se eliminó el cliente"))
}
