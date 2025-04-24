package service

import (
	"backend/internal/model"
	"github.com/google/uuid"
)

type SurveyService interface {
	CreateSurvey(newSurvey model.CreateSurveyRequest) (uuid.UUID, error)
	CreateQuestion(newQuestion []model.CreateQuestionRequest) error
	CreateAnswer(userId uuid.UUID, newAnswer model.CreateAnswerRequest) error
	GetSurveysDoneByUser(userId uuid.UUID) ([]*model.SurveyResponse, []*model.SurveyResponse, error)
	GetAnswers(surveyID uuid.UUID, questionID int) ([]model.AnswerResponse, error)
	DeleteAnswerByUser(userId uuid.UUID, req model.CreateAnswerRequest) error
	DeleteSurveyByID(userId uuid.UUID, surveyID uuid.UUID) error
}
