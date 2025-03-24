package delivery

import (
	"context"
	"go/rest/internal/entity"

	"github.com/google/uuid"
)

type IUseCase interface {
	CreateTask(context.Context, entity.Task) (uuid.UUID, error)
	GetTask(context.Context) ([]entity.Task, error)
	DeleteTask(context.Context, string) error
	UpdateTask(context.Context, entity.Task) error
}
