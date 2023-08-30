package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskInput struct {
	UserId      primitive.ObjectID `bson:"user_id"`
	WorkspaceId primitive.ObjectID `json:"workspace_id" bson:"workspace_id" binding:"required"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Content     string             `json:"content" bson:"content" binding:"required"`
	DateCreated time.Time          `bson:"date_created"`
}

type TaskOutput struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id" binding:"required"`
	WorkspaceId primitive.ObjectID `json:"workspace_id" bson:"workspace_id" binding:"required"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Content     string             `json:"content" bson:"content" binding:"required"`
	DateCreated time.Time          `json:"date_created" bson:"date_created" binding:"required"`
}

func (taskIn *TaskInput) InitTask(userId primitive.ObjectID) {
	taskIn.UserId = userId
	taskIn.DateCreated = time.Now()
}
