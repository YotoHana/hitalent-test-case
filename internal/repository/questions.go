package repository

import (
	"context"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"gorm.io/gorm"
)

type QuestionRepository interface {
	Create(ctx context.Context, q *models.Question) error
	GetByID(ctx context.Context, id int) (*models.Question, error)
	GetAll(ctx context.Context) (*[]models.Question, error)
	Delete(ctx context.Context, id int) error
}

type questionRepo struct {
	db *gorm.DB
}

func (r *questionRepo) Create(ctx context.Context, q *models.Question) error {
	return r.db.WithContext(ctx).Create(&q).Error
}

func (r *questionRepo) GetByID(ctx context.Context, id int) (*models.Question, error) {
	var question models.Question

	err := r.db.WithContext(ctx).First(&question, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &question, err
}

func (r *questionRepo) GetAll(ctx context.Context) (*[]models.Question, error) {
	var questions []models.Question

	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&questions).Error

	return &questions, err
}

func (r *questionRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Question{}, id).Error
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepo{db: db}
}
