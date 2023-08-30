package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceInput struct {
	OwnerId     primitive.ObjectID   `bson:"owner_id"`
	UsersId     []primitive.ObjectID `json:"users_id" bson:"users_id"`
	Title       string               `json:"title" bson:"title" binding:"required"`
	DateCreated time.Time            `bson:"date_created"`
}

type WorkspaceOutput struct {
	Id          primitive.ObjectID   `json:"id" bson:"_id" binding:"required"`
	OwnerId     primitive.ObjectID   `json:"owner_id" bson:"owner_id" binding:"required"`
	UsersId     []primitive.ObjectID `json:"users_id" bson:"users_id" binding:"required"`
	Title       string               `json:"title" bson:"title" binding:"required"`
	DateCreated time.Time            `json:"date_created" bson:"date_created" binding:"required"`
}

func (wsIn *WorkspaceInput) InitWorkSpace(userId primitive.ObjectID) {
	wsIn.OwnerId = userId
	wsIn.UsersId = append(wsIn.UsersId, userId)
	wsIn.DateCreated = time.Now()
}
