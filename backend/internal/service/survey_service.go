package service

import (
	"backend/internal/model"
	repo "backend/internal/repository"
	"backend/internal/util"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type SurveyService interface {
	CreateSurvey(newSurvey model.CreateSurveyRequest) error
	CreateQuestion(newQuestion model.CreateQuestionRequest) error
	CreateAnswer(userId uuid.UUID, newAnswer model.CreateAnswerRequest) error
	GetAllSurveys() ([]*model.SurveyResponse, error)
	GetSurveysDoneByUser(userId uuid.UUID) ([]*model.SurveyResponse, error)
	GetAnswers(surveyID uuid.UUID, questionID int) ([]model.AnswerResponse, error)
	DeleteAnswerByUser(userId uuid.UUID, req model.CreateAnswerRequest) error
}

type surveyService struct{}

func NewSurveyService() SurveyService {
	return &surveyService{}
}

func (s *surveyService) CreateSurvey(newSurvey model.CreateSurveyRequest) error {
	survey := &model.Survey{
		CreatedBy: newSurvey.UserID,
		Name:      newSurvey.Name,
	}
	if err := repo.CreateSurvey(survey); err != nil {
		log.Println("CreateSurvey error:", err)
		return fmt.Errorf("failed to create survey")
	}
	return nil
}

func (s *surveyService) CreateQuestion(newQuestion model.CreateQuestionRequest) error {
	question := &model.Question{
		Question: newQuestion.Question,
		SurveyID: newQuestion.SurveyID,
	}
	if err := repo.CreateQuestion(question); err != nil {
		log.Println("CreateQuestion error:", err)
		return fmt.Errorf("failed to create question")
	}
	return nil
}

func (s *surveyService) CreateAnswer(userId uuid.UUID, newAnswer model.CreateAnswerRequest) error {
	ans, err := util.DecodeAnswerFromBase64(newAnswer.Answer)
	if err != nil {
		log.Println("DecodeAnswerFromBase64 error:", err)
		return fmt.Errorf("invalid answer format")
	}
	answer := &model.Answer{
		UserID:     userId,
		SurveyID:   newAnswer.SurveyID,
		QuestionID: newAnswer.QuestionID,
		Answer:     ans,
	}
	if err := repo.CreateAnswer(answer); err != nil {
		log.Println("CreateAnswer error:", err)
		return fmt.Errorf("failed to submit answer")
	}
	return nil
}

func (s *surveyService) GetAllSurveys() ([]*model.SurveyResponse, error) {
	surveys, err := repo.GetAllSurveys()
	if err != nil {
		log.Println("GetAllSurveys error:", err)
		return nil, fmt.Errorf("could not retrieve surveys")
	}

	surveyResponses := transformDBtoOutput(surveys)

	return surveyResponses, nil
}

func (s *surveyService) GetSurveysDoneByUser(userId uuid.UUID) ([]*model.SurveyResponse, error) {
	surveys, err := repo.GetAllSurveysDoneByUser(userId)
	if err != nil {
		log.Println("GetAllSurveys error:", err)
		return nil, fmt.Errorf("could not retrieve surveys")
	}

	surveyResponses := transformDBtoOutput(surveys)

	return surveyResponses, nil
}

func (s *surveyService) GetAnswers(surveyID uuid.UUID, questionID int) ([]model.AnswerResponse, error) {
	answers, err := repo.GetAnswersBySurveyIDAndQuestionID(surveyID, questionID)
	if err != nil {
		log.Println("GetAnswersBySurveyIDAndQuestionID error:", err)
		return nil, fmt.Errorf("failed to retrieve answers")
	}

	var response []model.AnswerResponse
	for _, a := range answers {
		response = append(response, model.AnswerResponse{
			ID:         a.ID,
			UserID:     a.UserID,
			SurveyID:   a.SurveyID,
			QuestionID: a.QuestionID,
			Answer:     util.EncodeAnswerToBase64(a.Answer),
		})
	}

	return response, nil
}

func (s *surveyService) GetAnswersByUser(userID uuid.UUID) ([]model.AnswerResponse, error) {
	answers, err := repo.GetAnswersByUserID(userID)
	if err != nil {
		log.Println("GetAnswersByUserID error:", err)
		return nil, fmt.Errorf("failed to retrieve answers")
	}

	var response []model.AnswerResponse
	for _, a := range answers {
		response = append(response, model.AnswerResponse{
			ID:         a.ID,
			UserID:     a.UserID,
			SurveyID:   a.SurveyID,
			QuestionID: a.QuestionID,
			Answer:     util.EncodeAnswerToBase64(a.Answer),
		})
	}

	return response, nil
}

func (s *surveyService) DeleteAnswerByUser(userId uuid.UUID, req model.CreateAnswerRequest) error {
	err := repo.DeleteAnswer(userId, req.SurveyID, req.QuestionID)
	if err != nil {
		return fmt.Errorf("failed to delete answer")
	}
	return nil
}

func transformDBtoOutput(surveys []*model.Survey) []*model.SurveyResponse {
	// Prepare a slice to hold the transformed SurveyResponse objects
	var surveyResponses []*model.SurveyResponse
	for _, survey := range surveys {
		// Transform the Survey into SurveyResponse
		surveyResponse := &model.SurveyResponse{
			ID:          survey.ID,
			Name:        survey.Name,
			CreatedDate: survey.CreatedDate,
			CreatedBy:   survey.User.Name,
		}

		// Transform Questions into QuestionResponse
		for _, question := range survey.Questions {
			questionResponse := model.QuestionResponse{
				ID:          question.ID,
				Question:    question.Question,
				CreatedDate: question.CreatedDate,
				SurveyID:    survey.ID,
			}

			// Transform Answers into AnswerResponse
			for _, answer := range question.Answers {
				answerResponse := model.AnswerResponse{
					ID:         answer.ID,
					UserID:     answer.UserID,
					SurveyID:   answer.SurveyID,
					QuestionID: answer.QuestionID,
					Answer:     string(util.EncodeAnswerToBase64(answer.Answer)),
				}
				// Append the AnswerResponse to the Answers slice of the QuestionResponse
				questionResponse.Answers = append(questionResponse.Answers, answerResponse)
			}

			// Append the QuestionResponse to the Questions slice of the SurveyResponse
			surveyResponse.Questions = append(surveyResponse.Questions, questionResponse)
		}

		// Append the SurveyResponse to the final slice
		surveyResponses = append(surveyResponses, surveyResponse)
	}
	return surveyResponses
}
