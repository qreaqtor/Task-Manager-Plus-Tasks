package models

type User struct {
	Username  string `json:"username" bson:"username" binding:"required"`
	FirstName string `json:"firstName" bson:"firstName" binding:"required"`
	LastName  string `json:"lastName" bson:"lastName" binding:"required"`
}
