package domain

import "context"

type RoomUseCase interface {
	Bridger
	ListAllRooms(ctx context.Context) (Rooms, error)
	CreateRoom(ctx context.Context, form RoomForm) error
	ListUserRooms(ctx context.Context, userID string) (Rooms, error)
	GetRoomById(ctx context.Context, roomID string) (Room, error)
	ListRoomParticipants(ctx context.Context, roomID string) ([]User, error)
	GetUserRoom(ctx context.Context, roomID string) (Room, error)
}
