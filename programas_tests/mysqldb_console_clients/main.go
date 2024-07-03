package main

import (
	"mysql/01/config"
	"mysql/01/connection"
	"mysql/01/utils"
)

func main() {
	utils.PrintTitle()
	config.Init()
	connection.Connect()
	defer connection.CloseConnection()
	utils.Startup()
}
