package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-starter-kit/internal/models"
)

func HealthCheck(c echo.Context) error {
	var (
		response = models.NewResponse()
	)

	return c.JSON(http.StatusOK, response.WithMessage("OK"))
}
