package routing

import (
	"go/rest/internal/app/delivery"

	"github.com/gin-gonic/gin"
)

func APIrout(handler *delivery.Handler) {
	r := gin.Default()
	Router := r.Group("/task")
	{
		Router.GET("/", handler.GetTask)
		Router.POST("/", handler.AddTask)
		Router.PUT("/:id", handler.UpdateTask)
		Router.DELETE("/:id", handler.DeleteTask)
	}
	r.Run(":8080")
}
