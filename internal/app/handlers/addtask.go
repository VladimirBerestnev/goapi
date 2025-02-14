package handlers

import (
	"go/rest/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	newTask.ID = uuid.New()
	models.Tasklist[newTask.ID] = newTask
	ctx.JSON(http.StatusOK, models.Tasklist)
}
