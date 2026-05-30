package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Exam struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Questions []Question    `bson:"questions" json:"questions"`
	StartTime time.Time     `bson:"start_time" json:"start_time"`
	EndTime   time.Time     `bson:"end_time" json:"end_time"`
}