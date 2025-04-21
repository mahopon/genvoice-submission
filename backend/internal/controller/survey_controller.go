package controller

import (
	"backend/internal/service"
)

type SurveyController struct {
	SurveyService service.SurveyService
}

func NewSurveyController(s service.SurveyService) *SurveyController {
	return &SurveyController{s}
}
