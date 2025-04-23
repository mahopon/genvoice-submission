package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Survey struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Name        string    `json:"name" gorm:"not null"`
	CreatedDate time.Time `json:"created_date" gorm:"not null"`
	CreatedBy   uuid.UUID `json:"created_by" gorm:"type:uuid;not null"`
	User        User      `json:"user" gorm:"foreignKey:CreatedBy;references:ID"`

	Questions []Question `json:"questions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:SurveyID"`
}

type Question struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Question    string    `json:"question" gorm:"not null"`
	CreatedDate time.Time `json:"created_date" gorm:"not null"`

	SurveyID uuid.UUID `json:"survey_id" gorm:"type:uuid;not null;index"`
	Survey   Survey    `json:"survey" gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Answers []Answer `json:"answers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Answer struct {
	ID     int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index:idx_user_question_survey,unique"`

	SurveyID uuid.UUID `json:"survey_id" gorm:"type:uuid;not null;index:idx_user_question_survey,unique"`
	Survey   Survey    `json:"survey" gorm:"foreignKey:SurveyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	QuestionID int      `json:"question_id" gorm:"not null;index:idx_user_question_survey,unique"`
	Question   Question `json:"question" gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Answer []byte `json:"answer" gorm:"type:bytea"`
}

func (s *Survey) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	s.CreatedDate = time.Now()
	return nil
}

func (q *Question) BeforeCreate(tx *gorm.DB) error {
	q.CreatedDate = time.Now()
	return nil
}
