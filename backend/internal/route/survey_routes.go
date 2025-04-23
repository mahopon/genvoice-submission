package route

import (
	"backend/internal/controller"
	"backend/internal/middleware"
	"backend/internal/service"

	"github.com/labstack/echo/v4"
)

func RegisterSurveyRoutes(g *echo.Group) {
	surveyService := service.NewSurveyService()
	surveyController := controller.NewSurveyController(surveyService)

	// POST /api/survey
	g.POST("", surveyController.CreateSurvey, middleware.JWTMiddleware())
	// GET /api/survey
	g.GET("", surveyController.GetSurveysDone, middleware.JWTMiddleware())
	// GET /api/survey/:surveyId/:questionId/answer
	g.GET("/:surveyId/:questionId/answer", surveyController.GetAnswersOfSurveyQuestion, middleware.JWTMiddleware())
	// POST /api/survey/question
	g.POST("/question", surveyController.CreateQuestion, middleware.JWTMiddleware())
	// DELETE /api/survey/delete/:surveyId
	g.DELETE("/delete/:surveyId", surveyController.DeleteSurveyByID, middleware.JWTMiddleware())
	// POST /api/survey/answer
	g.POST("/answer", surveyController.CreateAnswer, middleware.JWTMiddleware())
}
