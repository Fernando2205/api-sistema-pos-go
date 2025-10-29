package main

import (
	"log"
	"sistema_pos_go/config"
	"sistema_pos_go/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//Conectar a base de datos
	config.ConnectDB()

	//Configurar rutas
	router := routes.SetupRoutes(config.DB)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Error al iniciar el servidor")
	}
}
