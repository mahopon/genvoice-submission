package repository

import (
	"backend/internal/model"
	"github.com/google/uuid"
)

// CreateSurvey creates a new survey in the database.
func CreateSurvey(survey *model.Survey) error {
	return db.Create(survey).Error
}

// CreateQuestion creates a new question for a survey in the database.
func CreateQuestion(question *model.Question) error {
	return db.Create(question).Error
}

// CreateAnswer creates a new answer for a question in the database.
func CreateAnswer(answer *model.Answer) error {
	return db.Create(answer).Error
}

// GetAllSurveys() retrieves all available surveys
func GetAllSurveys() ([]*model.Survey, error) {
	var surveys []*model.Survey

	// Preload related data (questions and answers) in one query
	if err := db.Preload("User").Preload("Questions").Find(&surveys).Error; err != nil {
		return nil, err
	}

	return surveys, nil
}

// GetAllSurveys() retrieves all available surveys
func GetSurvey(surveyId uuid.UUID) ([]*model.Survey, error) {
	var surveys []*model.Survey

	// Preload related data (questions and answers) in one query
	if err := db.Preload("User").Preload("Questions").Where("survey_id = ?", surveyId).Find(&surveys).Error; err != nil {
		return nil, err
	}

	return surveys, nil
}

// GetAnswersBySurveyIDAndQuestionID retrieves all answers for a specific survey and question.
func GetAnswersBySurveyIDAndQuestionID(surveyID uuid.UUID, questionID int) ([]model.Answer, error) {
	var answers []model.Answer
	if err := db.Where("survey_id = ? AND question_id = ?", surveyID, questionID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

// GetAnswersBySurveyIDAndQuestionID retrieves all answers for a specific survey and question for a specific user.
func GetAnswersBySurveyIDAndQuestionIDAndUserID(surveyID uuid.UUID, questionID int, userID uuid.UUID) ([]model.Answer, error) {
	var answers []model.Answer
	if err := db.Where("survey_id = ? AND question_id = ? AND user_id = ?", surveyID, questionID, userID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

// GetAnswersBySurveyIDAndQuestionID retrieves all answers for specific user.
func GetAnswersByUserID(userID uuid.UUID) ([]model.Answer, error) {
	var answers []model.Answer
	if err := db.Where("user_id = ?", userID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}
