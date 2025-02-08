package handlers

import (
	"github.com/wildanfaz/go-starter-kit/internal/services"
)

type UsersHandler struct {
	usersService services.UsersService
}

func NewUsersHandler(
	usersService services.UsersService,
) UsersHandler {
	return UsersHandler{
		usersService: usersService,
	}
}
