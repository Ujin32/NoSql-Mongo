package models

type User struct {
	Name        string      `bson:"name"`
	Sex         string      `bson:"sex"`
	Smartphone  *Smartphone `bson:"smartphone"`
	DateOfBirth int         `bson:"date_of_birth"`
}
