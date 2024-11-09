package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/project-app-inventaris/config"
	"github.com/project-app-inventaris/internal/app/delivery/controller"
	"github.com/project-app-inventaris/internal/app/delivery/middleware"
	"github.com/project-app-inventaris/internal/app/manager"
)

func SetupRouter(router *gin.Engine) error {

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	// User Controller
	userController := controller.NewUserController(repoManager.UserService(), repoManager.AuthService())

	// Category Controller
	categoryController := controller.NewCategoryController(repoManager.CategoryService())

	v1 := router.Group("/api/v1")
	{
		inventaris := v1.Group("/inventaris")
		{
			auth := inventaris.Group("/auth")
			{
				auth.POST("/register", userController.Registration)
				auth.POST("/login", userController.Login)
			}

			users := inventaris.Group("/users", middleware.AuthMiddleware())
			category := users.Group("/category")
			{
				category.POST("/", categoryController.CreateCategory)
				category.GET("/list", categoryController.FindAllCategory)
				category.GET("/:id", categoryController.FindCategory)
				category.PUT("/:id", categoryController.UpdateCategory)
				category.DELETE("/:id", categoryController.DeleteCategory)
			}

		}
	}

	return router.Run()

}
