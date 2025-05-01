package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	DueDate     string             `bson:"due_date,omitempty"`
	Status      string             `bson:"status,omitempty"`
}

type TaskRepository interface {
	CreateTask( c context.Context , task Task)  error
	GetTask( c context.Context , id primitive.ObjectID) (Task, error)
	GetAllTasks( c context.Context) ([]Task, error)
	UpdateTask( c context.Context , id primitive.ObjectID, task Task) (Task, error)
	DeleteTask( c context.Context , id primitive.ObjectID) error
}

type TaskUsecase interface {
	CreateTask( c context.Context , task Task)  error
	GetTask( c context.Context , id primitive.ObjectID) (Task, error)
	GetAllTasks( c context.Context) ([]Task, error)
	UpdateTask( c context.Context , id primitive.ObjectID, task Task) (Task, error)
	DeleteTask( c context.Context , id primitive.ObjectID) error
}

