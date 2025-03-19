package usecase

import (
	"go/rest/internal/entity"

	"github.com/google/uuid"
)

type UseCase struct {
	repository IDatabase
}

func New(repo IDatabase) *UseCase {
	return &UseCase{repository: repo}
}

func (uc *UseCase) CreateTask(task entity.Task) (uuid.UUID, error) {
	task.ID = uuid.New()
	return task.ID, uc.repository.Create(task)
}

func (uc *UseCase) GetTask() ([]entity.Task, error) {
	return uc.repository.Get()
}

func (uc *UseCase) DeleteTask(s string) error {
	return uc.repository.Delete(s)
}

func (uc *UseCase) UpdateTask(task entity.Task) error {
	return uc.repository.Update(task)
}
