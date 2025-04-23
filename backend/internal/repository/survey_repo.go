package repository

import (
	"backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

// CreateSurvey creates a new survey in the database.
func CreateSurvey(survey *model.Survey) error {
	return db.Create(survey).Error
}

// CreateQuestion creates a new question for a survey in the database.
func CreateQuestion(questions []*model.Question) error {
	if len(questions) == 0 {
		return nil
	}
	return db.Create(&questions).Error
}

// CreateAnswer creates a new answer for a question in the database.
func CreateAnswer(answer *model.Answer) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "question_id"}, {Name: "survey_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"answer"}),
	}).Create(answer).Error
}

// DeleteSurveyByID deletes a survey by its ID
func DeleteSurveyByID(surveyID uuid.UUID) error {
	// Delete the survey and its related data (if any)
	if err := db.Where("id = ?", surveyID).Delete(&model.Survey{}).Error; err != nil {
		return err
	}

	return nil
}

// GetAllSurveysDoneByUserSorted retrieves all surveys and categorizes them into:
// 1. Surveys created by the user
// 2. Surveys answered by the user (but not created by them)
// 3. Surveys neither created nor answered by the user
func GetAllSurveysDoneByUserSorted(userId uuid.UUID) (createdByUser []*model.Survey, otherSurveys []*model.Survey, err error) {
	var surveys []*model.Survey

	err = db.
		Preload("Questions.Answers").
		Preload("User").
		Preload("Questions").
		Find(&surveys).Error
	if err != nil {
		return nil, nil, err
	}

	var answered, unanswered []*model.Survey

	for _, survey := range surveys {
		if survey.CreatedBy == userId {
			createdByUser = append(createdByUser, survey)
			continue
		}

		userAnswered := false
		var filteredQuestions []model.Question

		for _, question := range survey.Questions {
			var userAnswers []model.Answer
			for _, answer := range question.Answers {
				if answer.UserID == userId {
					userAnswers = append(userAnswers, answer)
				}
			}

			if len(userAnswers) > 0 {
				userAnswered = true
				question.Answers = userAnswers
				filteredQuestions = append(filteredQuestions, question)
			}
		}

		if userAnswered {
			survey.Questions = filteredQuestions
			answered = append(answered, survey)
		} else {
			unanswered = append(unanswered, survey)
		}
	}

	otherSurveys = append(answered, unanswered...)

	return createdByUser, otherSurveys, nil
}

// GetAllSurveys() retrieves all available surveys
func GetSurvey(surveyId uuid.UUID) ([]*model.Survey, error) {
	var surveys []*model.Survey

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

// GetAnswersBySurveyIDAndQuestionID retrieves all answers for specific user.
func GetAnswersByUserID(userID uuid.UUID) ([]model.Answer, error) {
	var answers []model.Answer
	if err := db.Where("user_id = ?", userID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

// DeleteAnswer deletes a specific record by a specific user.
func DeleteAnswer(userID uuid.UUID, surveyID uuid.UUID, questionID int) error {
	return db.
		Where("user_id = ? AND survey_id = ? AND question_id = ?", userID, surveyID, questionID).
		Delete(&model.Answer{}).
		Error
}

// GetSurveyById gets a survey by its ID.
func GetSurveyById(surveyID uuid.UUID) (*model.Survey, error) {
	survey := &model.Survey{}
	if err := db.Preload("User").Where("id = ?", surveyID).First(survey).Error; err != nil {
		return nil, err
	}
	return survey, nil
}
