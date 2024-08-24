package domain

import "context"

type RoomUseCase interface {
	Bridger
	ListRooms(ctx context.Context, searchQuery string) (Rooms, error)
	CreateRoom(ctx context.Context, form RoomForm) error
	ListUserRooms(ctx context.Context, userID string) (Rooms, error)
	GetRoomById(ctx context.Context, roomID string) (Room, error)
	ListRoomParticipants(ctx context.Context, roomID string) ([]User, error)
	SearchRoom(ctx context.Context, searchQuery string) (Rooms, error)
	UpdateRoom(ctx context.Context, id string, roomForm RoomForm) error
	GetUserRoom(ctx context.Context, roomID string) (Room, error)
	DeleteUserRoom(ctx context.Context, roomID string) error
}
