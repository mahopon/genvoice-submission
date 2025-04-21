package route

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	userGroup := api.Group("/user")
	surveyGroup := api.Group("/survey")

	RegisterUserRoutes(userGroup)
	RegisterSurveyRoutes(surveyGroup)
}
