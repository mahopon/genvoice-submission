package route

import (
	"backend/internal/controller"
	"backend/internal/service"

	"github.com/labstack/echo/v4"
)

func RegisterSurveyRoutes(g *echo.Group) {
	surveyService := service.NewSurveyService()
	surveyController := controller.NewSurveyController(surveyService)

	// POST /api/survey
	g.POST("", surveyController.CreateSurvey)
	// GET /api/survey
	g.GET("", surveyController.GetAllSurveys)
	// GET /api/survey/:surveyId/:questionId/answer
	g.GET("/:surveyId/:questionId/answer", surveyController.GetAnswersOfSurveyQuestion)
	// POST /api/survey/question
	g.POST("/question", surveyController.CreateQuestion)

	// POST /api/survey/answer
	g.POST("/answer", surveyController.CreateAnswer)

}
