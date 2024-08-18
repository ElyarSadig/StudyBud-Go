package usecase

import (
	"context"
	"time"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"github.com/elyarsadig/studybud-go/pkg/utils"
)

type RoomUseCase struct {
	repositories map[string]domain.Bridger
	errHandler   errorHandler.Handler
	logger       logger.Logger
}

func NewRoom(errHandler errorHandler.Handler, logger logger.Logger, repositories ...domain.Bridger) domain.RoomUseCase {
	room := &RoomUseCase{
		repositories: make(map[string]domain.Bridger),
		errHandler:   errHandler,
		logger:       logger,
	}

	for _, repository := range repositories {
		switch repository.(type) {
		case domain.RoomRepository:
			room.repositories[configs.ROOMS_DB_NAME] = repository
		}
	}

	return room
}

func (u *RoomUseCase) None() {}

func (u *RoomUseCase) ListAllRooms(ctx context.Context) (domain.Rooms, error) {
	repo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	rooms, err := repo.ListAllRooms(ctx)
	if err != nil {
		return domain.Rooms{}, err
	}
	for i, room := range rooms.List {
		rooms.List[i].Since = utils.FormatDuration(time.Since(room.Created))
	}
	return rooms, nil
}