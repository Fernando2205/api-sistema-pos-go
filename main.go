package main

import (
	"log"
	"os"
	"sistema_pos_go/config"
	"sistema_pos_go/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env")
	}

	// Conectar a base de datos
	config.ConnectDB()

	// Configurar rutas
	router := routes.SetupRoutes(config.DB)

	// Obtener puerto del entorno o usar valor por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	serverAddr := "localhost:" + port
	log.Printf("Servidor iniciando en %s", serverAddr)

	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
