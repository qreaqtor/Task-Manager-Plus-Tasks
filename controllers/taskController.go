package controllers

import (
	"net/http"
	"task-manager-plus-tasks/models"
	"task-manager-plus-tasks/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Taskservice services.TaskService
}

func NewTaskController() TaskController {
	return TaskController{
		Taskservice: services.NewTaskService(),
	}
}

func (tc *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.Taskservice.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

// func (ts *TaskController) GetUserTasks(ctx *gin.Context) {
// 	userId := ctx.MustGet("userId").(primitive.ObjectID)

// }

func (tc *TaskController) RegisterTasksRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", tc.CreateTask)
}
