package usecase

import (
	"context"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
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
	return repo.ListAllMessages(ctx)
}
