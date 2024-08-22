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
		Order("created DESC").
		Find(&rooms.List).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Rooms{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return rooms, nil
}

func (r *RoomRepository) CreateRoom(ctx context.Context, room *domain.Room) error {
	err := r.db.WithContext(ctx).Create(&room).Error
	if err != nil {
		r.logger.Error(err.Error())
		return r.errHandler.New(http.StatusInternalServerError, "something went wrong")
	}
	return nil
}

func (r *RoomRepository) ListUserRooms(ctx context.Context, userID string) (domain.Rooms, error) {
	rooms := domain.Rooms{}
	err := r.db.WithContext(ctx).Model(&domain.Room{}).Where("host_id = ?", userID).Count(&rooms.Count).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Rooms{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	err = r.db.
		WithContext(ctx).
		Model(&domain.Room{}).
		Preload("Host").
		Preload("Topic").
		Where("host_id = ?", userID).
		Joins("LEFT JOIN room_participants ON room_participants.room_id = rooms.id").
		Select("rooms.*, COUNT(room_participants.id) as participants_count").
		Group("rooms.id").
		Order("created DESC").
		Find(&rooms.List).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Rooms{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return rooms, nil
}

func (r *RoomRepository) GetRoomById(ctx context.Context, roomID string) (domain.Room, error) {
	var tempRoom domain.Room
	err := r.db.WithContext(ctx).Model(&domain.Room{}).Preload("Host").Preload("Topic").Where("id = ?", roomID).First(&tempRoom).Error
	if err != nil {
		r.logger.Error(err.Error())
		return domain.Room{}, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return tempRoom, nil
}

func (r *RoomRepository) ListRoomParticipants(ctx context.Context, roomID string) ([]domain.RoomParticipant, error) {
	var users []domain.RoomParticipant
	err := r.db.WithContext(ctx).Model(&domain.RoomParticipant{}).Preload("User").Where("room_id = ?", roomID).Find(&users).Error
	if err != nil {
		return nil, r.errHandler.New(http.StatusInternalServerError, "something went wrong!")
	}
	return users, nil
}

func (r *RoomRepository) DeleteUserRoom(ctx context.Context, roomID, hostID string) error {
	err := r.db.WithContext(ctx).Model(&domain.Room{}).Where("id = ? AND host_id = ?", roomID, hostID).Delete(nil).Error
	if err != nil {
		r.logger.Error(err.Error())
		return r.errHandler.New(http.StatusInternalServerError, "something went wrong")
	}
	return nil
}
