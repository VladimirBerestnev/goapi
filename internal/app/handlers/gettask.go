package handlers

import (
	"go/rest/internal/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTask(ctx *gin.Context) {

	statusStr := ctx.Query("status")
	priorityStr := ctx.Query("priority")

	ctx.Header("Cache-Control", "public, max-age=3600")

	if statusStr == "" && priorityStr == "" {
		ctx.JSON(http.StatusOK, models.Tasklist)
		return
	}

	var tasks []models.Task
	for _, task := range models.Tasklist {
		if statusStr != "" {
			status, err := strconv.ParseBool(statusStr)
			if err != nil || task.Status != status {
				continue
			}
		}

		if priorityStr != "" && task.Priority != priorityStr {
			continue
		}
		tasks = append(tasks, task)
	}
	ctx.JSON(http.StatusOK, tasks)
}
