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

	//Inicializar servicios
	categoriaService := services.NewCategoriaService(categoriaRepo)

	//Inicilizar hanlders
	categoriaHandler := handlers.NewCategoriaHandler(categoriaService)

	//Grupo de rutas
	api := router.Group("/api")
	{
		//Rutas categoria
		categorias := api.Group("/categoria")
		{
			categorias.GET("/all", categoriaHandler.GetAll)
			categorias.GET("/:id", categoriaHandler.GetById)
			categorias.POST("/save", categoriaHandler.Create)
			categorias.PUT("update/:id", categoriaHandler.Update)
			categorias.DELETE("delete/:id", categoriaHandler.Delete)
			//categorias.PATCH("partial-update/:id", categoriaHandler.PartialUpdate)
		}
	}
	return router
}
