package controller

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/internal/util"
	"fmt"
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
	tokenString, err := ctx.Cookie("access_token")
	if err != nil {
		log.Printf("ERR: %v", err)
		return echo.ErrUnauthorized
	}

	claims, err := util.ValidateJWT(tokenString.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired access token"})
	}

	var req model.CreateSurveyRequest
	if err := ctx.Bind(&req); err != nil {
		log.Printf("ERR: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	req.UserID, _ = uuid.Parse(claims.Subject)

	// Manual validation
	if req.Name == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Name cannot be empty"})
	}

	id, err := c.SurveyService.CreateSurvey(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"survey_id": id})
}

func (c *SurveyController) CreateQuestion(ctx echo.Context) error {
	var req []model.CreateQuestionRequest
	if err := ctx.Bind(&req); err != nil {
		log.Printf("ERR: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	for _, question := range req {
		if question.SurveyID == uuid.Nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": fmt.Sprintln("SurveyID cannot be empty")})
		}
		if question.Question == "" {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": fmt.Sprintln("Question cannot be empty")})
		}
	}

	err := c.SurveyService.CreateQuestion(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Question(s) created successfully"})
}

func (c *SurveyController) CreateAnswer(ctx echo.Context) error {
	var req []model.CreateAnswerRequest
	if err := ctx.Bind(&req); err != nil {
		log.Printf("ERR: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	// Extract the JWT access_token from cookies
	tokenString, err := ctx.Cookie("access_token")
	if err != nil {
		log.Printf("ERR: %v", err)
		return echo.ErrUnauthorized
	}

	// Validate and extract the UserID from the JWT token using ValidateJWT function
	claims, err := util.ValidateJWT(tokenString.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired access token"})
	}

	// Get the UserID from the claims (assuming it's stored in "sub")
	userID := claims.Subject
	if userID == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	// Loop through each request item to perform validation and fill the UserID
	for _, answerReq := range req {
		// Check if SurveyID is valid
		if answerReq.SurveyID == uuid.Nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		// Check if QuestionID is valid
		if answerReq.QuestionID == 0 {
			return ctx.NoContent(http.StatusBadRequest)
		}

		parsedUUID, err := uuid.Parse(userID)
		if err != nil {
			return echo.ErrBadRequest
		}

		// If answer is empty, delete instead
		if answerReq.Answer == "" {
			err = c.SurveyService.DeleteAnswerByUser(parsedUUID, answerReq)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
			}
		} else {
			err = c.SurveyService.CreateAnswer(parsedUUID, answerReq)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
			}
		}

	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Answers sent successfully"})
}

func (c *SurveyController) GetSurveysDone(ctx echo.Context) error {

	tokenString, err := ctx.Cookie("access_token")
	if err != nil {
		log.Printf("ERR: %v", err)
		return echo.ErrUnauthorized
	}

	claims, err := util.ValidateJWT(tokenString.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired access token"})
	}

	userID := claims.Subject
	if userID == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	parsedUUID, _ := uuid.Parse(userID)

	userSurveys, otherSurveys, err := c.SurveyService.GetSurveysDoneByUser(parsedUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"user_made": userSurveys, "others_made": otherSurveys})
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

func (c *SurveyController) DeleteSurveyByID(ctx echo.Context) error {
	surveyIDParam := ctx.Param("surveyId")

	tokenString, err := ctx.Cookie("access_token")
	if err != nil {
		log.Printf("ERR: %v", err)
		return echo.ErrUnauthorized
	}

	claims, err := util.ValidateJWT(tokenString.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired access token"})
	}

	userID := claims.Subject
	if userID == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	surveyID, err := uuid.Parse(surveyIDParam)
	if err != nil {
		log.Printf("Error: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid survey ID"})
	}
	parsedUUID, err := uuid.Parse(userID)

	err = c.SurveyService.DeleteSurveyByID(parsedUUID, surveyID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}
	return nil
}
