package middleware

import (
	"net/http"

	"gitlab.com/labstack/echo"
)

type AppMiddleware struct {
	AppName string
}

func (am *AppMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Server", am.AppName)
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Response().Header().Set("Content-Type", "application/json")

		if c.Request().Method == "OPTIONS" {
			return c.String(http.StatusOK, "")
		}

		return next(c)
	}
}

func InitAppMiddleware(appName string) *AppMiddleware {
	return &AppMiddleware{
		AppName: appName,
	}
}
