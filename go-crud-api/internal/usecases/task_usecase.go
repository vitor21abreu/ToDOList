package usecases

import (
	"context"
	"go-crud-api/internal/entities"
	"go-crud-api/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error)
	GetTasks(ctx context.Context) ([]entities.Task, error)
	UpdateTask(ctx context.Context, id primitive.ObjectID, task *entities.Task) error
	DeleteTask(ctx context.Context, id primitive.ObjectID) error
}

type TaskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error) {
	return uc.repo.Create(ctx, task)
}

func (uc *TaskUseCase) GetTasks(ctx context.Context) ([]entities.Task, error) {
	return uc.repo.GetALL(ctx)
}

func (uc *TaskUseCase) UpdateTask(ctx context.Context, id primitive.ObjectID, task *entities.Task) error {
	return uc.repo.Update(ctx, id, task)
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
	return uc.repo.Delete(ctx, id)
}
