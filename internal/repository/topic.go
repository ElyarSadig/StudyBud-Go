package repository

import (
	"context"
	"net/http"

	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"gorm.io/gorm"
)

type TopicRepository struct {
	db         *gorm.DB
	errHandler errorHandler.Handler
	logger     logger.Logger
}

func NewTopic(db *gorm.DB, errHandler errorHandler.Handler, logger logger.Logger) domain.TopicRepository {
	return &TopicRepository{
		db:         db,
		errHandler: errHandler,
		logger:     logger,
	}
}

func (r *TopicRepository) None() {}

func (r *TopicRepository) ListAllTopics(ctx context.Context) (domain.Topics, error) {
	topics := domain.Topics{}
	err := r.db.WithContext(ctx).Model(domain.Topic{}).Count(&topics.Count).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Topics{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	err = r.db.
		WithContext(ctx).
		Table("topics").
		Select("topics.name, COUNT(rooms.id) as room_count").
		Joins("LEFT JOIN rooms ON rooms.topic_id = topics.id").
		Group("topics.name").
		Order("topics.name").
		Scan(&topics.List).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Topics{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return topics, nil
}

func (r *TopicRepository) SearchTopicByName(ctx context.Context, name string) (domain.Topics, error) {
	topics := domain.Topics{}
	err := r.db.WithContext(ctx).Model(domain.Topic{}).Count(&topics.Count).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Topics{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	err = r.db.
		WithContext(ctx).
		Table("topics").
		Select("topics.name, COUNT(rooms.id) as room_count").
		Joins("LEFT JOIN rooms ON rooms.topic_id = topics.id").
		Where("topics.name ILIKE ?", "%"+name+"%").
		Group("topics.name").
		Scan(&topics.List).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Topics{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return topics, nil
}
