package repository

import (
	"context"
	"net/http"

	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db         *gorm.DB
	errHandler errorHandler.Handler
	logger     logger.Logger
}

func NewMessage(db *gorm.DB, errHandler errorHandler.Handler, logger logger.Logger) domain.MessageRepository {
	return &MessageRepository{
		db:         db,
		errHandler: errHandler,
		logger:     logger,
	}
}

func (r *MessageRepository) None() {}

func (r *MessageRepository) ListAllMessages(ctx context.Context) (domain.Messages, error) {
	messages := domain.Messages{}
	err := r.db.WithContext(ctx).Model(domain.Message{}).Count(&messages.Count).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Messages{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	err = r.db.WithContext(ctx).Model(domain.Message{}).Preload("Room").Preload("User").Order("created DESC").Limit(5).Find(&messages.MessageList).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Messages{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return messages, nil
}

func (r *MessageRepository) Get(ctx context.Context, id string) (domain.Message, error) {
	var tempMessage domain.Message
	err := r.db.WithContext(ctx).Model(&domain.Message{}).Preload("User").Where("id = ?", id).First(&tempMessage).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Message{}, r.errHandler.New(http.StatusNotFound, "not found")
	}
	return tempMessage, nil
}

func (r *MessageRepository) Delete(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Model(&domain.Message{}).Delete("id = ?", id).Error
	if err != nil {
		r.logger.Error(err.Error())
		return r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return nil
}
