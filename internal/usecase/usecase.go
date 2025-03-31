package usecase

import (
	"context"
	"go/rest/internal/entity"
	"go/rest/internal/repository"

	"github.com/google/uuid"
)

type UseCase struct {
	repository repository.IDatabase
}

func New(repo repository.IDatabase) *UseCase {
	return &UseCase{repository: repo}
}

func (uc *UseCase) CreateTask(c context.Context, task entity.Task) (uuid.UUID, error) {
	task.ID = uuid.New()
	return task.ID, uc.repository.Create(c, task)
}

func (uc *UseCase) GetTask(c context.Context) ([]entity.Task, error) {
	return uc.repository.Get(c)
}

func (uc *UseCase) DeleteTask(c context.Context, s string) error {
	return uc.repository.Delete(c, s)
}

func (uc *UseCase) UpdateTask(c context.Context, task entity.Task) error {
	return uc.repository.Update(c, task)
}
