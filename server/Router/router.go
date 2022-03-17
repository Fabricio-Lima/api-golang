package routes

import (
	controllers "github.com/Fabricio-Lima/api-golang/controller"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		games := main.Group("games")
		{
			games.GET("/games/:id", controllers.ShowGame)
			games.GET("/", controllers.ShowGames)
			games.POST("/add", controllers.CreateGame)
			games.PUT("/update", controllers.UpdateGame)
			games.DELETE("/delete/:id", controllers.DeleteGame)

		}
	}

	return router
}
