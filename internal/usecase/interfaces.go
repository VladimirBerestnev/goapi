package usecase

import (
	"context"
	"go/rest/internal/entity"
)

type IDatabase interface {
	Create(context.Context, entity.Task) error
	Get(context.Context) ([]entity.Task, error)
	Delete(context.Context, string) error
	Update(context.Context, entity.Task) error
}
