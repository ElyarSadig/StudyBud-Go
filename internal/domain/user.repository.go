package domain

import "context"

type UserRepository interface {
	Bridger
	Create(ctx context.Context, obj *User) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}
