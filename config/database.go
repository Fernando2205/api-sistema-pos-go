package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	//DSN = Data Source Name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota", os.Getenv("DB_HOST"),
		os.Getenv("DSCBM_DB_USER"),
		os.Getenv("DSCBM_DB_PASSWORD"),
		os.Getenv("DSCBM_DB_NAME"),
		os.Getenv("DB_PORT"))

	//Conexión a base de datos
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Error al conectar a la base de datos")
	}

	fmt.Println("Conexión exitosa a base de datos")

}
