package routes

import (
	"sistema_pos_go/handlers"
	"sistema_pos_go/repositories"
	"sistema_pos_go/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configura todas las rutas de la aplicación
func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	//Inicializar repos
	categoriaRepo := repositories.NewCategoriaRepository(db)
	mesaRepo := repositories.NewMesaRepository(db)

	//Inicializar servicios
	categoriaService := services.NewCategoriaService(categoriaRepo)
	mesaService := services.NewMesaService(mesaRepo)

	//Inicilizar hanlders
	categoriaHandler := handlers.NewCategoriaHandler(categoriaService)
	mesaHandler := handlers.NewMesaHandler(mesaService)

	//Grupo de rutas
	api := router.Group("/api")
	{
		//Rutas categoria
		categorias := api.Group("/categorias")
		{
			categorias.GET("", categoriaHandler.GetAll)
			categorias.GET("/:id", categoriaHandler.GetById)
			categorias.POST("", categoriaHandler.Create)
			categorias.PUT("/:id", categoriaHandler.Update)
			categorias.DELETE("/:id", categoriaHandler.Delete)
		}
		//Rutas mesa
		mesas := api.Group("/mesas")
		{
			mesas.GET("", mesaHandler.GetAll)
			mesas.GET("/:id", mesaHandler.GetById)
			mesas.POST("", mesaHandler.Create)
			mesas.PUT("/:id", mesaHandler.Update)
			mesas.PATCH("/:id", mesaHandler.PartialUpdate)
			mesas.DELETE("/:id", mesaHandler.Delete)
		}

	}
	return router
}
