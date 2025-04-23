package controller

import (
	"backend/internal/model"
	"backend/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SurveyController struct {
	SurveyService service.SurveyService
}

func NewSurveyController(s service.SurveyService) *SurveyController {
	return &SurveyController{s}
}

func (c *SurveyController) CreateSurvey(ctx echo.Context) error {
	var req model.CreateSurveyRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	// Manual validation
	if req.UserID == uuid.Nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "UserID cannot be nil"})
	}
	if req.Name == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Name cannot be empty"})
	}

	err := c.SurveyService.CreateSurvey(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Survey created successfully"})
}

func (c *SurveyController) CreateQuestion(ctx echo.Context) error {
	var req model.CreateQuestionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	// Manual validation
	if req.SurveyID == uuid.Nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "SurveyID cannot be nil"})
	}
	if req.Question == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Question cannot be empty"})
	}

	err := c.SurveyService.CreateQuestion(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Question created successfully"})
}

func (c *SurveyController) CreateAnswer(ctx echo.Context) error {
	var req model.CreateAnswerRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	if req.SurveyID == uuid.Nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "SurveyID cannot be nil"})
	}
	if req.QuestionID == 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "QuestionID cannot be 0"})
	}
	if req.UserID == uuid.Nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "UserID cannot be nil"})
	}
	if req.Answer == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Answer cannot be empty"})
	}

	err := c.SurveyService.CreateAnswer(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Answer created successfully"})
}

func (c *SurveyController) AnswerQuestion(ctx echo.Context) error {
	var req model.CreateAnswerRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	if req.SurveyID == uuid.Nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "SurveyID cannot be nil"})
	}
	if req.QuestionID == 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "QuestionID cannot be 0"})
	}
	if req.UserID == uuid.Nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "UserID cannot be nil"})
	}
	if req.Answer == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Answer cannot be empty"})
	}

	err := c.SurveyService.CreateAnswer(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Answer recorded successfully"})
}

func (c *SurveyController) GetAllSurveys(ctx echo.Context) error {
	surveys, err := c.SurveyService.GetAllSurveys()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, surveys)
}

func (c *SurveyController) GetAnswersOfSurveyQuestion(ctx echo.Context) error {
	surveyIDParam := ctx.Param("surveyId")
	questionIdParam := ctx.Param("questionId")

	surveyID, err := uuid.Parse(surveyIDParam)
	if err != nil {
		log.Printf("Error: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid survey ID"})
	}
	questionId, err := strconv.Atoi(questionIdParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid question ID"})
	}

	answers, err := c.SurveyService.GetAnswers(surveyID, questionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, answers)
}
