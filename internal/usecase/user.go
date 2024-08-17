package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/bcrypt"
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
	err := u.validateUserRegisterForm(form)
	if err != nil {
		return "", err
	}
	repo := domain.Bridge[domain.UserRepository](configs.USERS_DB_NAME, u.repositories)
	hashedPassword, err := bcrypt.HashPassword(form.Password1)
	if err != nil {
		u.logger.Error(err.Error())
		return "", u.errHandler.New(http.StatusBadRequest, "something went wrong!")
	}
	user := domain.User{
		Name:     form.Name,
		Username: form.Username,
		Email:    form.Email,
		Password: hashedPassword,
	}
	_, err = repo.Create(ctx, &user)
	if err != nil {
		return "", err
	}
	key := u.generateDigitString(6)
	err = u.redis.Set(ctx, "session", time.Hour*1, key, user.Username)
	if err != nil {
		u.logger.Error(err.Error())
		return "", u.errHandler.New(http.StatusInternalServerError, "something went wrong")
	}
	return key, nil
}
