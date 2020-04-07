package router

import (
	"github.com/gin-gonic/gin"
	"rbac/app/controllers"
	"rbac/app/middlewares"
)

func InitRouter() {
	router := gin.Default()
	v1 := router.Group("v1")
	v1.Use(middlewares.Cors())
	{
		admin := v1.Group("admin")
		{
			admin.POST("/login", userControllers.Login)
			admin.Use(middlewares.JWTAuth(), middlewares.PerMission())
			{
				admin.GET("/user", userControllers.Index)
				admin.POST("/user", userControllers.Store)
				admin.PUT("/user/:id", userControllers.Update)
				admin.DELETE("/user/:id", userControllers.Delete)
			}

		}

	}

	router.Run(":8099")
}
