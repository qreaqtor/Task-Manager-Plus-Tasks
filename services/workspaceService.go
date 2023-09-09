package services

import (
	"context"
	"errors"
	"task-manager-plus-tasks/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkspaceService struct {
	ctx        *context.Context
	workspaces *mongo.Collection
}

func NewWorkspaceService() WorkspaceService {
	return WorkspaceService{
		workspaces: workspacesCollection,
		ctx:        &ctx,
	}
}

func (wss *WorkspaceService) CreateWorkspace(wsIn *models.WorkspaceInput) error {
	_, err := wss.workspaces.InsertOne(*wss.ctx, wsIn)
	return err
}

func (wss *WorkspaceService) GetUserWorkspaces(username string) (wspaces []models.WorkspaceOutput, err error) {
	filter := bson.M{"users.username": username}
	cursor, err := wss.workspaces.Find(*wss.ctx, filter)
	if err != nil {
		return
	}
	err = cursor.All(*wss.ctx, &wspaces)
	if err != nil {
		return
	}
	return
}

func (wss *WorkspaceService) DeleteWorkspace(workspace_id primitive.ObjectID) error {
	filter := bson.M{"_id": workspace_id}
	result, err := wss.workspaces.DeleteOne(*wss.ctx, filter)
	if err != err {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
