package handlers

import (
	"go/rest/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PutTask(ctx *gin.Context) {
	parsedID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}

	newTask, exists := models.Tasklist[parsedID]
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	models.Tasklist[parsedID] = newTask
	ctx.JSON(http.StatusOK, newTask)
}
