package domain

import (
	"context"
)

type UserUseCase interface {
	Bridger
	RegisterUser(ctx context.Context, form *UserRegisterForm) (string, error)
	Login(ctx context.Context, form *UserLoginForm) (string, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateInfo(ctx context.Context, obj *UpdateUser) (string, error)
}
