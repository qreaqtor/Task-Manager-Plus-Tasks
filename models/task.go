package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskInput struct {
	FromUser    string             `bson:"from_user"`
	ToUser      string             `json:"to_user" bson:"to_user" binding:"required"`
	WorkspaceId primitive.ObjectID `json:"workspace_id" bson:"workspace_id" binding:"required"`
	Content     string             `json:"content" bson:"content" binding:"required"`
	DateCreated time.Time          `bson:"date_created"`
}

type TaskOutput struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	FromUser    string             `json:"from_user" bson:"from_user" binding:"required"`
	ToUser      string             `json:"to_user" bson:"to_user" binding:"required"`
	WorkspaceId primitive.ObjectID `json:"workspace_id" bson:"workspace_id" binding:"required"`
	Content     string             `json:"content" bson:"content" binding:"required"`
	DateCreated time.Time          `json:"date_created" bson:"date_created" binding:"required"`
}

func (taskIn *TaskInput) InitTask(username string) {
	taskIn.FromUser = username
	taskIn.DateCreated = time.Now()
}
