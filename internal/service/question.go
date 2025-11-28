package service

import (
	"context"
	"errors"
	"time"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"github.com/YotoHana/hitalent-test-case/internal/repository"
)

type QuestionService interface {
	CreateQuestion(ctx context.Context, req *models.CreateQuestionRequest) error
	GetAllQuestions(ctx context.Context) (*[]models.Question, error)
}

type questionService struct {
	questionRepo repository.QuestionRepository
}

func (s *questionService) CreateQuestion(ctx context.Context, req *models.CreateQuestionRequest) error {
	if len(req.Text) < 3 {
		return errors.New("question text too short")
	}

	question := models.Question{
		Text: req.Text,
		CreatedAt: time.Now(),
	}

	err := s.questionRepo.Create(ctx, &question)
	if err != nil {
		return err
	}

	return nil
}

func (s *questionService) GetAllQuestions(ctx context.Context) (*[]models.Question, error) {
	questions, err := s.questionRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func NewQuestionService(repo repository.QuestionRepository) QuestionService {
	return &questionService{questionRepo: repo}
}