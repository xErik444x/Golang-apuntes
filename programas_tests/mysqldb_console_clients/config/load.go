package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig Config

func Init() {
	log.Println("Loading .env...")
	errorVars := godotenv.Load()
	if errorVars != nil {
		log.Panic("Error al intentar cargar archivo .env")
	}

	// Cargar variables de entorno
	AppConfig = Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
	}

	log.Println("Loading .env...OK")
}
