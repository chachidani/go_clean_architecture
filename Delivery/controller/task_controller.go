package controller

import (
	"go_clean_architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := tc.TaskUsecase.CreateTask(c, task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	task, err := tc.TaskUsecase.GetTask(c, objectId)
	if err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(http.StatusNotFound, domain.ErrorResponse{Message: "Task not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAllTasks(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var updateTask domain.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	updatedTask, err := tc.TaskUsecase.UpdateTask(c, objectId, updateTask)
	if err != nil {
		if err.Error() == "task not found" {
			c.IndentedJSON(http.StatusNotFound, domain.ErrorResponse{Message: "Task not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userRole := c.GetHeader("Role")
	err = tc.TaskUsecase.DeleteTask(c, objectId, userRole)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
