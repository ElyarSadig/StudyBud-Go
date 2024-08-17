package usecase

import (
	"context"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
)

type TopicUseCase struct {
	repositories map[string]domain.Bridger
	errHandler   errorHandler.Handler
	logger       logger.Logger
}

func NewTopic(errHandler errorHandler.Handler, logger logger.Logger, repositories ...domain.Bridger) domain.TopicUseCase {
	topic := &TopicUseCase{
		repositories: make(map[string]domain.Bridger),
		errHandler:   errHandler,
		logger:       logger,
	}

	for _, repository := range repositories {
		switch repository.(type) {
		case domain.TopicRepository:
			topic.repositories[configs.TOPICS_DB_NAME] = repository
		}
	}

	return topic
}

func (u *TopicUseCase) None() {}

func (u *TopicUseCase) ListAllTopics(ctx context.Context) (domain.Topics, error) {
	repo := domain.Bridge[domain.TopicRepository](configs.TOPICS_DB_NAME, u.repositories)
	return repo.ListAllTopics(ctx)
}

func (u *TopicUseCase) SearchTopicByName(ctx context.Context, name string) (domain.Topics, error) {
	repo := domain.Bridge[domain.TopicRepository](configs.TOPICS_DB_NAME, u.repositories)
	return repo.SearchTopicByName(ctx, name)
}
