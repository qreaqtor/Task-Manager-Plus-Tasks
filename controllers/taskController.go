package controllers

import (
	"net/http"
	"task-manager-plus-tasks/models"
	"task-manager-plus-tasks/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var taskIn models.TaskInput
	if err := ctx.ShouldBindJSON(&taskIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userId := ctx.MustGet("userId").(primitive.ObjectID)
	taskIn.InitTask(userId)
	err := tc.Taskservice.CreateTask(taskIn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (tc *TaskController) GetUserTasks(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(primitive.ObjectID)
	tasks, err := tc.Taskservice.GetUsersTasks(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (tc *TaskController) DeleteUserTask(ctx *gin.Context) {
	taskId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = tc.Taskservice.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succes"})
}

func (tc *TaskController) RegisterTasksRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", tc.CreateTask)
	rg.GET("/get/by-user", tc.GetUserTasks)
	rg.DELETE("/delete/:id", tc.DeleteUserTask)
}
