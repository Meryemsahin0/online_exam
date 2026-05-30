package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Question struct {
	ID            bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Content       string        `bson:"content" json:"content"`
	Options       []string      `bson:"options" json:"options"`
	CorrectAnswer string        `bson:"correct_answer" json:"correct_answer"`
}