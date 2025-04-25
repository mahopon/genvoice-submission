package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateSurveyRequest struct {
	SurveyID uuid.UUID // Injected
	UserID   uuid.UUID `json:"user_id"`
	Name     string    `json:"name"`
}

type CreateQuestionRequest struct {
	UserID   uuid.UUID // Injected
	SurveyID uuid.UUID `json:"survey_id"`
	Question string    `json:"question"`
}

type CreateAnswerRequest struct {
	SurveyID   uuid.UUID `json:"survey_id"`
	QuestionID int       `json:"question_id"`
	Answer     string    `json:"answer"`
}

type SurveyResponse struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	CreatedDate   time.Time          `json:"created_date"`
	CreatedBy     string             `json:"created_by"`
	CreatedByName string             `json:"created_by_name"`
	Questions     []QuestionResponse `json:"questions,omitempty"`
}

type QuestionResponse struct {
	ID          int              `json:"id"`
	Question    string           `json:"question"`
	CreatedDate time.Time        `json:"created_date"`
	SurveyID    uuid.UUID        `json:"survey_id"`
	Answers     []AnswerResponse `json:"answers,omitempty"`
}

type AnswerResponse struct {
	ID         int       `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	SurveyID   uuid.UUID `json:"survey_id"`
	QuestionID int       `json:"question_id"`
	Answer     string    `json:"answer"`
}
