package repository

import (
	"context"
	"net/http"

	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"gorm.io/gorm"
)

type RoomRepository struct {
	db         *gorm.DB
	errHandler errorHandler.Handler
	logger     logger.Logger
}

func NewRoom(db *gorm.DB, errHandler errorHandler.Handler, logger logger.Logger) domain.RoomRepository {
	return &RoomRepository{
		db:         db,
		errHandler: errHandler,
		logger:     logger,
	}
}

func (r *RoomRepository) None() {}

func (r *RoomRepository) ListAllRooms(ctx context.Context) (domain.Rooms, error) {
	rooms := domain.Rooms{}
	err := r.db.WithContext(ctx).Model(&domain.Room{}).Count(&rooms.Count).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Rooms{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	err = r.db.
		WithContext(ctx).
		Model(&domain.Room{}).
		Preload("Host").
		Preload("Topic").
		Joins("LEFT JOIN room_participants ON room_participants.room_id = rooms.id").
		Select("rooms.*, COUNT(room_participants.id) as participants_count").
		Group("rooms.id").
		Find(&rooms.List).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Rooms{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return rooms, nil
}
