package delivery

import (
	"go/rest/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	uc IUseCase
}

func New(ucase IUseCase) *Handler {
	return &Handler{uc: ucase}
}

func (h *Handler) AddTask(ctx *gin.Context) {
	var newTask entity.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := h.uc.CreateTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetTask(ctx *gin.Context) {
	statusStr := ctx.Query("status")
	priorityStr := ctx.Query("priority")

	ctx.Header("Cache-Control", "public, max-age=3600")

	tasks, err := h.uc.GetTask()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if statusStr == "" && priorityStr == "" {
		ctx.JSON(http.StatusOK, tasks)
		return
	}

	var taskEnv []entity.Task
	for _, task := range tasks {
		if statusStr != "" {
			status, err := strconv.ParseBool(statusStr)
			if err != nil || task.Status != status {
				continue
			}
		}
		if priorityStr != "" && task.Priority != priorityStr {
			continue
		}
		taskEnv = append(taskEnv, task)
	}
	ctx.JSON(http.StatusOK, taskEnv)
}

func (h *Handler) DeleteTask(ctx *gin.Context) {

	parsedID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error ID"})
		return
	}
	stringID := parsedID.String()

	if err = h.uc.DeleteTask(stringID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": stringID})
}

func (h *Handler) UpdateTask(ctx *gin.Context) {

	parsedID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error ID"})
		return
	}
	var newTask entity.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	newTask.ID = parsedID
	err = h.uc.UpdateTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, newTask)
}
