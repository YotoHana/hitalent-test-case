package repository

import (
	"context"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(ctx context.Context, q *models.Answer) error
	GetByQuestionID(ctx context.Context, questionID int) (*[]models.Answer, error)
	GetByID(ctx context.Context, id int) (*models.Answer, error)
	Delete(ctx context.Context, id int) error
}

type answerRepo struct {
	db *gorm.DB
}

func (r *answerRepo) Create(ctx context.Context, q *models.Answer) error {
	return r.db.WithContext(ctx).Create(&q).Error
}

func (r *answerRepo) GetByQuestionID(ctx context.Context, questionID int) (*[]models.Answer, error) {
	var answers []models.Answer

	err := r.db.WithContext(ctx).Find(&answers, questionID).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &answers, err
}

func (r *answerRepo) GetByID(ctx context.Context, id int) (*models.Answer, error) {
	var answer models.Answer
	err := r.db.WithContext(ctx).First(&answer, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &answer, err
}

func (r *answerRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Answer{}, id).Error
}

func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepo{db: db}
}
