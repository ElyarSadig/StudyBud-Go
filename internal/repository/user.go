package repository

import (
	"context"

	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"gorm.io/gorm"
)

type UserRepository struct {
	db         *gorm.DB
	errHandler errorHandler.Handler
	logger     logger.Logger
}

func NewUser(db *gorm.DB, errHandler errorHandler.Handler, logger logger.Logger) domain.UserRepository {
	return &UserRepository{
		db:         db,
		errHandler: errHandler,
		logger:     logger,
	}
}

func (r *UserRepository) None() {}

func (r *UserRepository) Create(ctx context.Context, obj *domain.User) (domain.User, error) {
	panic("Not Implemented")
}
