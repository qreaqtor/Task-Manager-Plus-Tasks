package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkspaceInput struct {
	Title       string    `json:"title" bson:"title" binding:"required"`
	Users       []string  `json:"users" bson:"users"`
	DateCreated time.Time `bson:"date_created"`
}

type WorkspaceOutput struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Users       []string           `json:"users" bson:"users" binding:"required"`
	DateCreated time.Time          `json:"date_created" bson:"date_created" binding:"required"`
}

func (wsIn *WorkspaceInput) InitWorkSpace(username string) {
	wsIn.Users = append(wsIn.Users, username)
	wsIn.DateCreated = time.Now()
}
