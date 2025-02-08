package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-starter-kit/configs"
	"github.com/wildanfaz/go-starter-kit/internal/handlers"
	"github.com/wildanfaz/go-starter-kit/internal/repositories"
	"github.com/wildanfaz/go-starter-kit/internal/services"
)

func InitEchoRouter() {
	// configs
	config := configs.InitConfig()
	dbMySql := configs.InitMySql(config.DatabaseDSN)

	// repositories
	usersRepository := repositories.NewUsersRepository(dbMySql)

	// services
	usersService := services.NewUsersService(usersRepository, config)

	// handlers
	usersHandler := handlers.NewUsersHandler(usersService)

	e := echo.New()

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/health", handlers.HealthCheck)

	UsersRouter(apiV1, usersHandler)

	e.Logger.Fatal(e.Start(config.AppPort))
}
