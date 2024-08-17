package domain

import (
	"context"
)

type UserUseCase interface {
	Bridger
	RegisterUser(ctx context.Context, form *UserRegisterForm) (string, error)
	Login(ctx context.Context, form *UserLoginForm) (string, error)
}
