package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	//DSN = Data Source Name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		os.Getenv("DB_HOST"),
		os.Getenv("DSCBM_DB_USER"),
		os.Getenv("DSCBM_DB_PASSWORD"),
		os.Getenv("DSCBM_DB_NAME"),
		os.Getenv("DSCBM_DB_PORT"))

	//Conexión a base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	fmt.Println("Conexión exitosa a base de datos")
	return db, nil
}
