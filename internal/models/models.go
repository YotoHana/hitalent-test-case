package models

import "time"

type Question struct {
	ID int `gorm:"primaryKey;column:id;type:serial" json:"id"`
	Text string `gorm:"column:text;type:text;not null" json:"text"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Answer struct {
	ID int `gorm:"primaryKey;column:id;type:serial" json:"id"`
	QuestionID int `gorm:"column:question_id;type:integer;not null" json:"question_id"`
	UserID string `gorm:"column:user_id;type:varchar(255);not null" json:"user_id"`
	Text string `gorm:"column:text;type:text;not null" json:"text"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

type DetailQuestion struct {
	ID int `json:"id"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Answers []Answer `json:"answers"`
}