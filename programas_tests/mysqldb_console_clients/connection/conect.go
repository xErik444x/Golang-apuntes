package connection

import (
	"database/sql"
	"log"
	"mysql/01/config"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() {
	conf := config.AppConfig
	con, err := sql.Open("mysql", conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+":"+conf.DBPort+")"+"/"+conf.DBName)

	if err != nil {
		log.Panic(err)
	}
	Db = con
	log.Println("Conectado a la base de datos")

}

func CloseConnection() {
	Db.Close()
}
