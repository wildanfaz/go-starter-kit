package middlewares

import "github.com/labstack/echo/v4"

func BearerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: implement middleware
		return next(c)
	}
}
