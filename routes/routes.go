package routes

import (
	"sistema_pos_go/handlers"
	"sistema_pos_go/middleware"
	"sistema_pos_go/repositories"
	"sistema_pos_go/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configura todas las rutas de la aplicación
func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Middlewares globales
	router.Use(middleware.RecoveryHandler())
	router.Use(middleware.ErrorHandler())

	// Handler para rutas no encontradas (404)
	router.NoRoute(middleware.NoRouteHandler)

	//Inicializar repos
	categoriaRepo := repositories.NewCategoriaRepository(db)
	mesaRepo := repositories.NewMesaRepository(db)
	empleadoRepo := repositories.NewEmpleadoRepository(db)

	//Inicializar servicios
	categoriaService := services.NewCategoriaService(categoriaRepo)
	mesaService := services.NewMesaService(mesaRepo)
	empleadoService := services.NewEmpleadoService(empleadoRepo)

	//Inicilizar hanlders
	categoriaHandler := handlers.NewCategoriaHandler(categoriaService)
	mesaHandler := handlers.NewMesaHandler(mesaService)
	empleadoHandler := handlers.NewEmpleadoHandler(empleadoService)

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
		//Rutas empleado
		empleados := api.Group("/empleados")
		{
			empleados.GET("", empleadoHandler.GetAll)
			empleados.GET("/:id", empleadoHandler.GetById)
			empleados.POST("", empleadoHandler.Create)
			empleados.PUT("/:id", empleadoHandler.Update)
			empleados.PATCH("/:id", empleadoHandler.PartialUpdate)
			empleados.DELETE("/:id", empleadoHandler.Delete)
		}

	}
	return router
}
