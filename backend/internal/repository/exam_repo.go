package repository

import (
	"context"
	"online-exam-app/internal/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ExamRepository interface {
	CreateExam(ctx context.Context, exam *domain.Exam) error
	GetAllExams(ctx context.Context) ([]domain.Exam, error)
	SaveResult(ctx context.Context, result *domain.ExamResult) error
	GetResult(ctx context.Context, examID bson.ObjectID, studentID string) (*domain.ExamResult, error)
}

type mongoExamRepo struct {
	examColl   *mongo.Collection
	resultColl *mongo.Collection
}

func NewMongoExamRepository(db *mongo.Database) ExamRepository {
	return &mongoExamRepo{
		examColl:   db.Collection("exams"),
		resultColl: db.Collection("results"),
	}
}

func (r *mongoExamRepo) CreateExam(ctx context.Context, exam *domain.Exam) error {
	exam.ID = bson.NewObjectID()
	_, err := r.examColl.InsertOne(ctx, exam)
	return err
}

func (r *mongoExamRepo) GetAllExams(ctx context.Context) ([]domain.Exam, error) {
	var exams []domain.Exam
	cursor, err := r.examColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &exams); err != nil {
		return nil, err
	}
	return exams, nil
}

func (r *mongoExamRepo) SaveResult(ctx context.Context, result *domain.ExamResult) error {
	result.ID = bson.NewObjectID()
	_, err := r.resultColl.InsertOne(ctx, result)
	return err
}

func (r *mongoExamRepo) GetResult(ctx context.Context, examID bson.ObjectID, studentID string) (*domain.ExamResult, error) {
	var result domain.ExamResult
	filter := bson.M{"exam_id": examID, "student_id": studentID}
	
	err := r.resultColl.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}