package repository

import (
	"context"
	"online-exam-app/internal/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuestionRepository interface {
	Create(ctx context.Context, q *domain.Question) error
	GetByIDs(ctx context.Context, ids []bson.ObjectID) ([]domain.Question, error)
}

type mongoQuestionRepo struct {
	collection *mongo.Collection
}

func NewMongoQuestionRepository(db *mongo.Database) QuestionRepository {
	return &mongoQuestionRepo{collection: db.Collection("questions")}
}

func (r *mongoQuestionRepo) Create(ctx context.Context, q *domain.Question) error {
	q.ID = bson.NewObjectID()
	_, err := r.collection.InsertOne(ctx, q)
	return err
}

func (r *mongoQuestionRepo) GetByIDs(ctx context.Context, ids []bson.ObjectID) ([]domain.Question, error) {
	var questions []domain.Question
	filter := bson.M{"_id": bson.M{"$in": ids}}
	
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &questions); err != nil {
		return nil, err
	}
	return questions, nil
}