package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"github.com/elyarsadig/studybud-go/pkg/utils"
)

type MessageUseCase struct {
	repositories map[string]domain.Bridger
	errHandler   errorHandler.Handler
	logger       logger.Logger
}

func NewMessage(errHandler errorHandler.Handler, logger logger.Logger, repositories ...domain.Bridger) domain.MessageUseCase {
	m := &MessageUseCase{
		repositories: make(map[string]domain.Bridger),
		errHandler:   errHandler,
		logger:       logger,
	}

	for _, repository := range repositories {
		switch repository.(type) {
		case domain.MessageRepository:
			m.repositories[configs.MESSAGES_DB_NAME] = repository
		}
	}

	return m
}

func (u *MessageUseCase) None() {}

func (u *MessageUseCase) ListAllMessages(ctx context.Context) (domain.Messages, error) {
	repo := domain.Bridge[domain.MessageRepository](configs.MESSAGES_DB_NAME, u.repositories)
	messages, err := repo.ListAllMessages(ctx)
	if err != nil {
		return domain.Messages{}, err
	}
	for i, message := range messages.MessageList {
		messages.MessageList[i].Since = utils.FormatDuration(time.Since(message.Created))
	}
	return messages, nil
}

func (u *MessageUseCase) GetUserMessage(ctx context.Context, id string) (domain.Message, error) {
	repo := domain.Bridge[domain.MessageRepository](configs.MESSAGES_DB_NAME, u.repositories)
	sessionValue := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	message, err := repo.Get(ctx, id)
	if err != nil {
		return domain.Message{}, err
	}
	if message.User.Username != sessionValue.Username {
		return domain.Message{}, u.errHandler.New(http.StatusForbidden, "forbidden!")
	}
	return message, nil
}

func (u *MessageUseCase) Delete(ctx context.Context, id string) error {
	sessionValue := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	repo := domain.Bridge[domain.MessageRepository](configs.MESSAGES_DB_NAME, u.repositories)
	message, err := repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if message.User.Username != sessionValue.Username {
		return u.errHandler.New(http.StatusUnauthorized, "forbidden!")
	}
	return repo.Delete(ctx, id)
}

func (u *MessageUseCase) ListUserMessages(ctx context.Context, userID string) (domain.Messages, error) {
	repo := domain.Bridge[domain.MessageRepository](configs.MESSAGES_DB_NAME, u.repositories)
	messages, err := repo.ListUserMessages(ctx, userID)
	if err != nil {
		return domain.Messages{}, err
	}
	for i, message := range messages.MessageList {
		messages.MessageList[i].Since = utils.FormatDuration(time.Since(message.Created))
	}
	return messages, nil
}

func (u *MessageUseCase) ListRoomMessages(ctx context.Context, roomID string) (domain.Messages, error) {
	repo := domain.Bridge[domain.MessageRepository](configs.MESSAGES_DB_NAME, u.repositories)
	messages, err := repo.ListRoomMessages(ctx, roomID)
	if err != nil {
		return domain.Messages{}, err
	}
	for i, message := range messages.MessageList {
		messages.MessageList[i].Since = utils.FormatDuration(time.Since(message.Created))
	}
	return messages, nil
}