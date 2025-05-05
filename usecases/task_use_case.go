package usecases

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_clean_architecture/domain"
	"time"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	ContextTimeout time.Duration
}


func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		ContextTimeout: timeout,
	}
}

// CreateTask implements domain.TaskUsecase.
func (t *taskUsecase) CreateTask(c context.Context, task domain.Task) error {
	ctx, cancel := context.WithTimeout(c, t.ContextTimeout)
	defer cancel()

	return t.taskRepository.CreateTask(ctx, task)
}

// DeleteTask implements domain.TaskUsecase.
func (t *taskUsecase) DeleteTask(c context.Context, id primitive.ObjectID, userRole string) error {
	ctx, cancel := context.WithTimeout(c, t.ContextTimeout)
	defer cancel()

	return t.taskRepository.DeleteTask(ctx, id, userRole)
}

// GetAllTasks implements domain.TaskUsecase.
func (t *taskUsecase) GetAllTasks(c context.Context) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.ContextTimeout)
	defer cancel()

	return t.taskRepository.GetAllTasks(ctx)
}

// GetTask implements domain.TaskUsecase.
func (t *taskUsecase) GetTask(c context.Context, id primitive.ObjectID) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.ContextTimeout)
	defer cancel()

	return t.taskRepository.GetTask(ctx, id)
}

// UpdateTask implements domain.TaskUsecase.
func (t *taskUsecase) UpdateTask(c context.Context, id primitive.ObjectID, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, t.ContextTimeout)
	defer cancel()

	return t.taskRepository.UpdateTask(ctx, id, task)
}


