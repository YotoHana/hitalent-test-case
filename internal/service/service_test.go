package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/YotoHana/hitalent-test-case/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) Create(ctx context.Context, q *models.Question) error {
	args := m.Called(ctx, q)
	q.ID = 1
	return args.Error(0)
}

func (m *MockQuestionRepository) GetByID(ctx context.Context, id int) (*models.Question, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionRepository) GetAll(ctx context.Context) (*[]models.Question, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.Question), args.Error(1)
}

func (m *MockQuestionRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockAnswerRepository struct {
	mock.Mock
}

func (m *MockAnswerRepository) Create(ctx context.Context, a *models.Answer) error {
	args := m.Called(ctx, a)
	a.ID = 1
	return args.Error(0)
}

func (m *MockAnswerRepository) GetByQuestionID(ctx context.Context, questionID int) (*[]models.Answer, error) {
	args := m.Called(ctx, questionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.Answer), args.Error(1)
}

func (m *MockAnswerRepository) GetByID(ctx context.Context, id int) (*models.Answer, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Answer), args.Error(1)
}

func (m *MockAnswerRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateQuestion_Success(t *testing.T) {
	mockQuestionRepo := new(MockQuestionRepository)
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewQuestionService(mockQuestionRepo, mockAnswerRepo)

	ctx := context.Background()
	req := &models.CreateQuestionRequest{
		Text: "What is unit testing?",
	}

	mockQuestionRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Question")).Return(nil)

	result, err := service.CreateQuestion(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Text, result.Text)
	assert.False(t, result.CreatedAt.IsZero())
	assert.False(t, result.ID == 0)

	mockQuestionRepo.AssertExpectations(t)
}

func TestCreateQuestion_TextTooShort(t *testing.T) {
	mockQuestionRepo := new(MockQuestionRepository)
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewQuestionService(mockQuestionRepo, mockAnswerRepo)

	ctx := context.Background()
	req := &models.CreateQuestionRequest{
		Text: "A?",
	}

	result, err := service.CreateQuestion(ctx, req)

	assert.Error(t, err)
	assert.Equal(t, "question text too short", err.Error())
	assert.Nil(t, result)

	mockQuestionRepo.AssertNotCalled(t, "Create")
}

func TestGetQuestionByID_Success(t *testing.T) {
	mockQuestionRepo := new(MockQuestionRepository)
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewQuestionService(mockQuestionRepo, mockAnswerRepo)

	ctx := context.Background()
	questionID := 1

	expectedAnswers := &[]models.Answer{
		{
			ID:         1,
			QuestionID: questionID,
			UserID:     "good user",
			Text:       "It`s so good",
			CreatedAt:  time.Now(),
		},
		{
			ID:         2,
			QuestionID: questionID,
			UserID:     "bad user",
			Text:       "It`s so bad",
			CreatedAt:  time.Now(),
		},
	}

	expectedQuestion := &models.Question{
		ID:        questionID,
		Text:      "What is unit testing?",
		CreatedAt: time.Now(),
	}

	expectedDetailQuestion := &models.DetailQuestion{
		ID:        expectedQuestion.ID,
		Text:      expectedQuestion.Text,
		CreatedAt: expectedQuestion.CreatedAt,
		Answers:   *expectedAnswers,
	}

	mockQuestionRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(expectedQuestion, nil)
	mockAnswerRepo.On("GetByQuestionID", mock.Anything, mock.AnythingOfType("int")).Return(expectedAnswers, nil)

	result, err := service.GetQuestionByID(ctx, questionID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedDetailQuestion, result)
	assert.Len(t, result.Answers, 2)

	mockAnswerRepo.AssertExpectations(t)
	mockQuestionRepo.AssertExpectations(t)
}

func TestGetQuestionByID_NotFound(t *testing.T) {
	mockQuestionRepo := new(MockQuestionRepository)
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewQuestionService(mockQuestionRepo, mockAnswerRepo)

	ctx := context.Background()
	questionID := 1
	errorNotFound := errors.New("Not Found")

	mockQuestionRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(nil, errorNotFound)

	result, err := service.GetQuestionByID(ctx, questionID)

	assert.Error(t, err)
	assert.Equal(t, errorNotFound, err)
	assert.Nil(t, result)

	mockQuestionRepo.AssertExpectations(t)
	mockAnswerRepo.AssertNotCalled(t, "GetByQuestionID")
}

func TestGetByQuestionID_NoAnswers(t *testing.T) {
	mockQuestionRepo := new(MockQuestionRepository)
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewQuestionService(mockQuestionRepo, mockAnswerRepo)

	ctx := context.Background()
	questionID := 1

	expectedQuestion := &models.Question{
		ID:        questionID,
		Text:      "What is unit testing?",
		CreatedAt: time.Now(),
	}

	errorNoAnswers := errors.New("Answers not found")

	mockQuestionRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(expectedQuestion, nil)
	mockAnswerRepo.On("GetByQuestionID", mock.Anything, mock.AnythingOfType("int")).Return(nil, errorNoAnswers)

	result, err := service.GetQuestionByID(ctx, questionID)

	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.Nil(t, result.Answers)
	assert.Equal(t, errorNoAnswers, err)

	mockQuestionRepo.AssertExpectations(t)
	mockAnswerRepo.AssertExpectations(t)
}

func TestCreateAnswer_Success(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewAnswerService(mockAnswerRepo)

	ctx := context.Background()
	req := models.CreateAnswerRequest{
		Text:   "Unit testing best",
		UserID: "user1",
	}
	questionID := 1

	mockAnswerRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Answer")).Return(nil)

	result, err := service.CreateAnswer(ctx, req, questionID)

	assert.NoError(t, err)
	assert.Equal(t, req.Text, result.Text)
	assert.NotNil(t, result)
	assert.False(t, result.CreatedAt.IsZero())
	assert.False(t, result.ID == 0)

	mockAnswerRepo.AssertExpectations(t)
}

func TestCreateAnswer_TextTooShort(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewAnswerService(mockAnswerRepo)

	ctx := context.Background()
	req := models.CreateAnswerRequest{
		UserID: "user1",
		Text:   "Y",
	}
	questionID := 1

	result, err := service.CreateAnswer(ctx, req, questionID)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "answer text too short", err.Error())

	mockAnswerRepo.AssertNotCalled(t, "Create")
}

func TestCreateAnswer_NotFound(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepository)

	service := NewAnswerService(mockAnswerRepo)

	ctx := context.Background()
	req := models.CreateAnswerRequest{
		UserID: "user1",
		Text:   "Unit testing best",
	}
	questionID := 1

	errorNotFound := errors.New("answer not found")

	mockAnswerRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.Answer")).Return(errorNotFound)

	result, err := service.CreateAnswer(ctx, req, questionID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errorNotFound, err)

	mockAnswerRepo.AssertExpectations(t)
}
