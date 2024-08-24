package repository

import (
	"context"
	"net/http"

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
	var tempUser domain.User
	r.db.WithContext(ctx).Where("email = ?", obj.Email).Find(&tempUser)
	if tempUser.Email != "" {
		return domain.User{}, r.errHandler.New(http.StatusConflict, "email already in use")
	}
	result := r.db.WithContext(ctx).Create(obj)
	err := result.Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.User{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return *obj, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var tempUser domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&tempUser).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.User{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return tempUser, nil
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) error {
	err := r.db.Model(&domain.User{}).WithContext(ctx).Where("email = ?", user.Email).Updates(user).Error
	if err != nil {
		r.logger.Error(err.Error())
		return r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id string) (domain.User, error) {
	var tempUser domain.User
	err := r.db.Model(&domain.User{}).WithContext(ctx).Where("id = ?", id).First(&tempUser).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.User{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return tempUser, nil
}
