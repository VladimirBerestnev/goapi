package handlers

import (
	"go/rest/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteTask(ctx *gin.Context) {

	parsedID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	_, exists := models.Tasklist[parsedID]
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
	delete(models.Tasklist, parsedID)
}
