package usecase

import (
	"context"
	"net/http"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/bcrypt"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	redispkg "github.com/elyarsadig/studybud-go/pkg/redis"
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
		IsActive: true,
	}
	_, err = repo.Create(ctx, &user)
	if err != nil {
		return "", err
	}
	sessionValue := domain.SessionValue{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	return u.setSession(ctx, sessionValue)
}

func (u *UserUseCase) Login(ctx context.Context, form *domain.UserLoginForm) (string, error) {
	repo := domain.Bridge[domain.UserRepository](configs.USERS_DB_NAME, u.repositories)
	user, err := repo.FindUserByEmail(ctx, form.Email)
	if err != nil {
		return "", err
	}
	ok := bcrypt.CheckPasswordHash(form.Password, user.Password)
	if !ok {
		return "", u.errHandler.New(http.StatusBadRequest, "invalid credentials try again!")
	}
	sessionValue := domain.SessionValue{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
	return u.setSession(ctx, sessionValue)
}
