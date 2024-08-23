package domain

import "context"

type RoomUseCase interface {
	Bridger
	ListAllRooms(ctx context.Context) (Rooms, error)
	CreateRoom(ctx context.Context, form RoomForm) error
	ListUserRooms(ctx context.Context, userID string) (Rooms, error)
	GetRoomById(ctx context.Context, roomID string) (Room, error)
	ListRoomParticipants(ctx context.Context, roomID string) ([]User, error)
	UpdateRoom(ctx context.Context, id string, roomForm RoomForm) error
	GetUserRoom(ctx context.Context, roomID string) (Room, error)
	DeleteUserRoom(ctx context.Context, roomID string) error
}
