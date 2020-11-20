package model

type Student struct {
	Id      int64  `bson:"id"`
	ClassId int64  `bson:"score,omitempty"`
	Name    string `bson:"name,omitempty"`
}
