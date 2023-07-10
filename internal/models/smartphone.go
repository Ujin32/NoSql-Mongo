package models

type Smartphone struct {
	Model  string `bson:"model"`
	Vendor string `bson:"vendor"`
}
