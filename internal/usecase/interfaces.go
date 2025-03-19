package usecase

import (
	"go/rest/internal/entity"
)

type IDatabase interface {
	Create(task entity.Task) error
	Get() ([]entity.Task, error)
	Delete(string) error
	Update(task entity.Task) error
}
