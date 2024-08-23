package usecase

import (
	"context"
	"net/http"
	"strconv"
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
		case domain.TopicRepository:
			room.repositories[configs.TOPICS_DB_NAME] = repository
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

func (u *RoomUseCase) CreateRoom(ctx context.Context, form domain.RoomForm) error {
	sessionValue := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	topicRepo := domain.Bridge[domain.TopicRepository](configs.TOPICS_DB_NAME, u.repositories)
	topic := domain.Topic{Name: form.TopicName}
	err := topicRepo.CreateTopicIfNotExists(ctx, &topic)
	if err != nil {
		return err
	}
	roomRepo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	room := domain.Room{
		Name:        form.Name,
		TopicID:     topic.ID,
		HostID:      uint(sessionValue.ID),
		Description: form.Description,
	}
	return roomRepo.CreateRoom(ctx, &room)
}

func (u *RoomUseCase) ListUserRooms(ctx context.Context, userID string) (domain.Rooms, error) {
	repo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	rooms, err := repo.ListUserRooms(ctx, userID)
	if err != nil {
		return domain.Rooms{}, err
	}
	for i, room := range rooms.List {
		rooms.List[i].Since = utils.FormatDuration(time.Since(room.Created))
	}
	return rooms, nil
}

func (u *RoomUseCase) GetRoomById(ctx context.Context, roomID string) (domain.Room, error) {
	repo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	room, err := repo.GetRoomById(ctx, roomID)
	if err != nil {
		return domain.Room{}, err
	}
	room.Since = utils.FormatDuration(time.Since(room.Created))
	return room, nil
}

func (u *RoomUseCase) ListRoomParticipants(ctx context.Context, roomID string) ([]domain.User, error) {
	repo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	participants, err := repo.ListRoomParticipants(ctx, roomID)
	if err != nil {
		return nil, err
	}
	users := make([]domain.User, 0, len(participants))
	for _, participant := range participants {
		users = append(users, participant.User)
	}
	return users, nil
}

func (u *RoomUseCase) GetUserRoom(ctx context.Context, roomID string) (domain.Room, error) {
	sv := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	repo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	room, err := repo.GetRoomById(ctx, roomID)
	if err != nil {
		return domain.Room{}, err
	}
	if room.Host.Username != sv.Username {
		return domain.Room{}, u.errHandler.New(http.StatusForbidden, "forbidden!")
	}
	return room, nil
}

func (u *RoomUseCase) DeleteUserRoom(ctx context.Context, roomID string) error {
	sv := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	hostID := strconv.Itoa(sv.ID)
	repo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	return repo.DeleteUserRoom(ctx, roomID, hostID)
}

func (u *RoomUseCase) UpdateRoom(ctx context.Context, id string, roomForm domain.RoomForm) error {
	repo := domain.Bridge[domain.RoomRepository](configs.ROOMS_DB_NAME, u.repositories)
	topicRepo := domain.Bridge[domain.TopicRepository](configs.TOPICS_DB_NAME, u.repositories)
	room, err := repo.GetRoomById(ctx, id)
	if err != nil {
		return err
	}
	topic := &domain.Topic{Name: roomForm.TopicName}
	err = topicRepo.CreateTopicIfNotExists(ctx, topic)
	if err != nil {
		return err
	}
	room.TopicID = topic.ID
	room.Name = roomForm.Name
	room.Description = roomForm.Description
	return repo.UpdateRoom(ctx, room)
}
