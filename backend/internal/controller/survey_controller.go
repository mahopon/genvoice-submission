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
	var req model.CreateSurveyRequest
	var err error
	if err := ctx.Bind(&req); err != nil {
		log.Printf("ERR: %v", err)
		return ctx.JSON(http.StatusForbidden, echo.Map{"message": "Invalid request format"})
	}

	req.UserID, err = util.GetUUIDFromContext(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	// Manual validation
	if req.Name == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Name cannot be empty"})
	}

	surveyId, err := c.SurveyService.CreateSurvey(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"survey_id": surveyId})
}

func (c *SurveyController) CreateQuestion(ctx echo.Context) error {
	var req []model.CreateQuestionRequest
	if err := ctx.Bind(&req); err != nil {
		log.Printf("ERR: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request format"})
	}

	for _, question := range req {
		if question.SurveyID == uuid.Nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"message": fmt.Sprintln("SurveyID cannot be empty")})
		}
		if question.Question == "" {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"message": fmt.Sprintln("Question cannot be empty")})
		}
	}

	err := c.SurveyService.CreateQuestion(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Question(s) created successfully"})
}

func (c *SurveyController) CreateAnswer(ctx echo.Context) error {
	var req []model.CreateAnswerRequest
	var err error
	if err := ctx.Bind(&req); err != nil {
		log.Printf("ERR: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request format"})
	}

	parsedUUID, err := util.GetUUIDFromContext(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user"})
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

		// If answer is empty, delete instead
		if answerReq.Answer == "" {
			err = c.SurveyService.DeleteAnswerByUser(parsedUUID, &answerReq)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
			}
		} else {
			err = c.SurveyService.CreateAnswer(parsedUUID, &answerReq)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
			}
		}

	}

	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Answers sent successfully"})
}

func (c *SurveyController) GetSurveysDone(ctx echo.Context) error {
	parsedUUID, err := util.GetUUIDFromContext(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user"})
	}

	userSurveys, otherSurveys, err := c.SurveyService.GetSurveysDoneByUser(parsedUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"user_made": userSurveys, "others_made": otherSurveys})
}

func (c *SurveyController) GetAnswersOfSurveyQuestion(ctx echo.Context) error {
	surveyIDParam := ctx.Param("surveyId")
	questionIdParam := ctx.Param("questionId")

	surveyID, err := uuid.Parse(surveyIDParam)
	if err != nil {
		log.Printf("Error: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid survey ID"})
	}
	questionId, err := strconv.Atoi(questionIdParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid question ID"})
	}

	answers, err := c.SurveyService.GetAnswers(surveyID, questionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, answers)
}

func (c *SurveyController) DeleteSurveyByID(ctx echo.Context) error {
	surveyIDParam := ctx.Param("surveyId")

	surveyID, err := uuid.Parse(surveyIDParam)
	if err != nil {
		log.Printf("Error: %v", err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid survey ID"})
	}

	parsedUUID, err := util.GetUUIDFromContext(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user"})
	}

	err = c.SurveyService.DeleteSurveyByID(parsedUUID, surveyID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": err})
	}
	return nil
}
