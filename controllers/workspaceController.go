package controllers

import (
	"net/http"
	"task-manager-plus-tasks/models"
	"task-manager-plus-tasks/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceController struct {
	workSpaceService   services.WorkspaceService
	deleteTasksHandler func(workspaceId primitive.ObjectID) error
}

func NewWorkSpaceController(tasksDeleteHandler func(workspaceId primitive.ObjectID) error) WorkspaceController {
	return WorkspaceController{
		workSpaceService:   services.NewWorkspaceService(),
		deleteTasksHandler: tasksDeleteHandler,
	}
}

func (wsc *WorkspaceController) createWorkspace(ctx *gin.Context) {
	var wsIn *models.WorkspaceInput
	if err := ctx.ShouldBindJSON(&wsIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	wsIn.InitWorkspace()
	err := wsc.workSpaceService.CreateWorkspace(wsIn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (wsc *WorkspaceController) getUserWorkspaces(ctx *gin.Context) {
	username := ctx.Param("username")
	wspaces, err := wsc.workSpaceService.GetUserWorkspaces(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": wspaces})
}

func (wsc *WorkspaceController) deleteWorkspace(ctx *gin.Context) {
	workspaceId, err := primitive.ObjectIDFromHex(ctx.Param("workspaceId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = wsc.deleteTasksHandler(workspaceId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = wsc.workSpaceService.DeleteWorkspace(workspaceId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (wsc *WorkspaceController) addUserToWorkspace(ctx *gin.Context) {
	workspaceId, err := primitive.ObjectIDFromHex(ctx.Param("workspaceId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var user models.User
	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = wsc.workSpaceService.AddUserToWorkspace(workspaceId, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (wsc *WorkspaceController) deleteUserFromWorkspace(ctx *gin.Context) {
	workspaceId, err := primitive.ObjectIDFromHex(ctx.Param("workspaceId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	username := ctx.Param("username")
	err = wsc.workSpaceService.DeleteUserFromWorkSpace(username, workspaceId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (wsc *WorkspaceController) RegisterWorkspaceRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", wsc.createWorkspace)
	rg.PATCH("/add/:workspaceId", wsc.addUserToWorkspace)
	rg.PATCH("/remove/:username/:workspaceId", wsc.deleteUserFromWorkspace)
	rg.GET("/get/all/:username", wsc.getUserWorkspaces)
	rg.DELETE("/delete/:workspaceId", wsc.deleteWorkspace)
}
