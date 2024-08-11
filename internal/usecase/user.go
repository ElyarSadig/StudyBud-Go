package usecase

import (
	"context"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	redispkg "github.com/elyarsadig/studybud-go/pkg/redisPkg"
)

type UserUseCase struct {
	repositories map[string]domain.Bridger
	errHandler   errorHandler.Handler
	redis        *redispkg.Redis
	logger       logger.Logger
}

func NewUser(errHandler errorHandler.Handler, redis *redispkg.Redis, logger logger.Logger, repositories ...domain.Bridger) domain.UserUseCase {
	user := &UserUseCase{
		repositories: make(map[string]domain.Bridger),
		errHandler:   errHandler,
		redis:        redis,
		logger:       logger,
	}

	for _, repository := range repositories {
		switch repository.(type) {
		case domain.UserRepository:
			user.repositories[configs.USERS_DB_NAME] = repository
		}
	}

	return user
}

func (u *UserUseCase) None() {}

func (u *UserUseCase) RegisterUser(ctx context.Context, form *domain.UserRegisterForm) (string, error) {
	panic("Not Implemented")
}
