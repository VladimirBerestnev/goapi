package delivery

import (
	"go/rest/internal/entity"

	"github.com/google/uuid"
)

type IUseCase interface {
	CreateTask(task entity.Task) (uuid.UUID, error)
	GetTask() ([]entity.Task, error)
	DeleteTask(string) error
	UpdateTask(task entity.Task) error
}
