package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username  string `json:"username" bson:"username" binding:"required"`
	FirstName string `json:"firstName" bson:"firstName" binding:"required"`
	LastName  string `json:"lastName" bson:"lastName" binding:"required"`
}

type WorkspaceInput struct {
	Title       string    `json:"title" bson:"title" binding:"required"`
	Owner       string    `json:"owner" bson:"owner" binding:"required"`
	Users       []User    `json:"users" bson:"users" binding:"required"`
	DateCreated time.Time `bson:"date_created"`
}

type WorkspaceOutput struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Owner       string             `json:"owner" bson:"owner" binding:"required"`
	Users       []User             `json:"users" bson:"users" binding:"required"`
	DateCreated time.Time          `json:"date_created" bson:"date_created" binding:"required"`
}

func (wsIn *WorkspaceInput) InitWorkspace() {
	wsIn.DateCreated = time.Now()
}
