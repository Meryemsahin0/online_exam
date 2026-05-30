package service

import (
	"context"
	"errors"
	"time"
	"online-exam-app/internal/domain"
	"online-exam-app/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StudentExamDTO struct {
	ID        bson.ObjectID `json:"id"`
	Title     string        `json:"title"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	IsActive  bool          `json:"is_active"`
	Score     *float64      `json:"score"`
}

type ExamService struct {
	questionRepo repository.QuestionRepository
	examRepo     repository.ExamRepository
}

func NewExamService(qRepo repository.QuestionRepository, eRepo repository.ExamRepository) *ExamService {
	return &ExamService{
		questionRepo: qRepo,
		examRepo:     eRepo,
	}
}

func (s *ExamService) AddQuestionToPool(ctx context.Context, content string, options []string, correct string) error {
	if content == "" || len(options) < 2 || correct == "" {
		return errors.New("geçersiz soru verisi")
	}
	q := &domain.Question{
		Content:       content,
		Options:       options,
		CorrectAnswer: correct,
	}
	return s.questionRepo.Create(ctx, q)
}

func (s *ExamService) CreateExam(ctx context.Context, title string, questionIDs []string, start time.Time, end time.Time) error {
	if start.After(end) {
		return errors.New("başlangıç tarihi bitiş tarihinden sonra olamaz")
	}

	var oids []bson.ObjectID
	for _, idStr := range questionIDs {
		oid, err := bson.ObjectIDFromHex(idStr)
		if err != nil {
			return errors.New("geçersiz soru ID'si")
		}
		oids = append(oids, oid)
	}

	questions, err := s.questionRepo.GetByIDs(ctx, oids)
	if err != nil {
		return err
	}

	exam := &domain.Exam{
		Title:     title,
		Questions: questions,
		StartTime: start,
		EndTime:   end,
	}

	return s.examRepo.CreateExam(ctx, exam)
}

func (s *ExamService) ListExamsForStudent(ctx context.Context, studentID string) ([]StudentExamDTO, error) {
	exams, err := s.examRepo.GetAllExams(ctx)
	if err != nil {
		return nil, err
	}

	var list []StudentExamDTO
	now := time.Now()

	for _, exam := range exams {
		result, err := s.examRepo.GetResult(ctx, exam.ID, studentID)
		var score *float64
		if err == nil && result != nil && result.IsGraded {
			score = &result.Score
		}

		isActive := now.After(exam.StartTime) && now.Before(exam.EndTime)

		list = append(list, StudentExamDTO{
			ID:        exam.ID,
			Title:     exam.Title,
			StartTime: exam.StartTime,
			EndTime:   exam.EndTime,
			IsActive:  isActive,
			Score:     score,
		})
	}

	return list, nil
}