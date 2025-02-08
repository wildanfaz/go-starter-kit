package services

import (
	"github.com/wildanfaz/go-starter-kit/configs"
	"github.com/wildanfaz/go-starter-kit/internal/repositories"
)

type ImplementUsersService struct {
	users  repositories.UsersReporitory
	config *configs.Config
}

type UsersService interface {
}

func NewUsersService(
	users repositories.UsersReporitory,
	config *configs.Config,
) UsersService {
	return &ImplementUsersService{
		users:  users,
		config: config,
	}
}
