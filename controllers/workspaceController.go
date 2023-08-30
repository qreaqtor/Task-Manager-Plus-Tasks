package controllers

import (
	"net/http"
	"task-manager-plus-tasks/models"
	"task-manager-plus-tasks/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceController struct {
	workSpaceService services.WorkspaceService
}

func NewWorkSpaceController() WorkspaceController {
	return WorkspaceController{
		workSpaceService: services.NewWorkspaceService(),
	}
}

func (wsc *WorkspaceController) createWorkspace(ctx *gin.Context) {
	var wsIn *models.WorkspaceInput
	if err := ctx.ShouldBindJSON(&wsIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userId := ctx.MustGet("userId").(primitive.ObjectID)
	wsIn.InitWorkSpace(userId)
	err := wsc.workSpaceService.CreateWorkspace(wsIn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (wsc *WorkspaceController) getUserWorkspaces(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(primitive.ObjectID)
	wspaces, err := wsc.workSpaceService.GetUserWorkspaces(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": wspaces})
}

func (wsc *WorkspaceController) deleteWorkspace(ctx *gin.Context) {
	workspace_id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = wsc.workSpaceService.DeleteWorkspace(workspace_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (wsc *WorkspaceController) RegisterWorkspaceRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", wsc.createWorkspace)
	rg.GET("/get/by-user", wsc.getUserWorkspaces)
	rg.DELETE("/delete/:id", wsc.deleteWorkspace)
}
