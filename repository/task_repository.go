package repository

import (
	"context"
	"fmt"
	"go_clean_architecture/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(database mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   database,
		collection: collection,
	}
}

// CreateTask implements domain.TaskRepository.
func (t *taskRepository) CreateTask(c context.Context, task domain.Task) error {
	collection := t.database.Collection(t.collection)
	_, err := collection.InsertOne(c, task)
	return err
}

// DeleteTask implements domain.TaskRepository.
func (t *taskRepository) DeleteTask(c context.Context, id primitive.ObjectID) error {
	collection := t.database.Collection(t.collection)
	_, err := collection.DeleteOne(c, bson.M{"_id": id})
	return err
}

// GetAllTasks implements domain.TaskRepository.
func (t *taskRepository) GetAllTasks(c context.Context) ([]domain.Task, error) {
	collection := t.database.Collection(t.collection)

	var tasks []domain.Task

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTask implements domain.TaskRepository.
func (t *taskRepository) GetTask(c context.Context, id primitive.ObjectID) (domain.Task, error) {
	collection := t.database.Collection(t.collection)

	var task domain.Task
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, fmt.Errorf("task not found")
		}
		return domain.Task{}, err
	}
	return task, nil
}

// UpdateTask implements domain.TaskRepository.
func (t *taskRepository) UpdateTask(c context.Context, id primitive.ObjectID, task domain.Task) (domain.Task, error) {
	collection := t.database.Collection(t.collection)

	var updatedTask domain.Task
	err := collection.FindOneAndUpdate(
		c,
		bson.M{"_id": id},
		bson.M{"$set": task},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedTask)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, fmt.Errorf("task not found")
		}
		return domain.Task{}, err
	}

	return updatedTask, nil
}
