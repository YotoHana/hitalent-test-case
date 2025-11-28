package service

import (
	"context"
	"errors"
	"time"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"github.com/YotoHana/hitalent-test-case/internal/repository"
)

type AnswerService interface {
	CreateAnswer(ctx context.Context, req models.CreateAnswerRequest, questionID int) error
}

type answerService struct {
	answerRepo repository.AnswerRepository
}

func (s *answerService) CreateAnswer(ctx context.Context, req models.CreateAnswerRequest, questionID int,) error {
	if len(req.Text) < 3 {
		return errors.New("answer text too short")
	}

	answer := models.Answer{
		QuestionID: questionID,
		UserID: req.UserID,
		Text: req.Text,
		CreatedAt: time.Now(),
	}

	return s.answerRepo.Create(ctx, &answer)
}

func NewAnswerService(answerRepo repository.AnswerRepository) AnswerService {
	return &answerService{
		answerRepo: answerRepo,
	}
}