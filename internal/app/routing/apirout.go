package routing

import (
	"go/rest/internal/app/files"
	"go/rest/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func APIrout() {

	r := gin.Default()
	Router := r.Group("/task")
	{
		Router.GET("/", handlers.GetTask)
		Router.POST("/", handlers.AddTask)
		Router.PUT("/:id", handlers.PutTask)
		Router.DELETE("/:id", handlers.DeleteTask)
		Router.POST("/save", files.SaveInFile)
	}
	r.Run(":8080")
}
