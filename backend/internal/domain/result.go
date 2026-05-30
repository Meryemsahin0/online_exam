package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type ExamResult struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	ExamID    bson.ObjectID `bson:"exam_id" json:"exam_id"`
	StudentID string        `bson:"student_id" json:"student_id"`
	Score     float64       `bson:"score" json:"score"`
	IsGraded  bool          `bson:"is_graded" json:"is_graded"`
}