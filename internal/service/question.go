package service

import (
	"context"
	"errors"
	"time"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"github.com/YotoHana/hitalent-test-case/internal/repository"
)

type QuestionService interface {
	CreateQuestion(ctx context.Context, req *models.CreateQuestionRequest) (*models.Question, error)
	GetAllQuestions(ctx context.Context) (*[]models.Question, error)
	GetQuestionByID(ctx context.Context, id int) (*models.DetailQuestion, error)
	DeleteQuestionByID(ctx context.Context, id int) error
}

type questionService struct {
	questionRepo repository.QuestionRepository
	answerRepo   repository.AnswerRepository
}

func (s *questionService) CreateQuestion(ctx context.Context, req *models.CreateQuestionRequest) (*models.Question, error) {
	if len(req.Text) < 3 {
		return nil, errors.New("question text too short")
	}

	question := models.Question{
		Text:      req.Text,
		CreatedAt: time.Now(),
	}

	err := s.questionRepo.Create(ctx, &question)
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (s *questionService) GetAllQuestions(ctx context.Context) (*[]models.Question, error) {
	questions, err := s.questionRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (s *questionService) GetQuestionByID(ctx context.Context, id int) (*models.DetailQuestion, error) {
	question, err := s.questionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	answers, err := s.answerRepo.GetByQuestionID(ctx, question.ID)
	if err != nil {
		return &models.DetailQuestion{
			ID:        question.ID,
			Text:      question.Text,
			CreatedAt: question.CreatedAt,
			Answers:   nil,
		}, err
	}

	return &models.DetailQuestion{
		ID:        question.ID,
		Text:      question.Text,
		CreatedAt: question.CreatedAt,
		Answers:   *answers,
	}, nil
}

func (s *questionService) DeleteQuestionByID(ctx context.Context, id int) error {
	return s.questionRepo.Delete(ctx, id)
}

func NewQuestionService(
	questionRepo repository.QuestionRepository,
	answerRepo repository.AnswerRepository,
) QuestionService {
	return &questionService{
		questionRepo: questionRepo,
		answerRepo:   answerRepo,
	}
}
