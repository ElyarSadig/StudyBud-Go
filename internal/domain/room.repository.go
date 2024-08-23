package domain

import "context"

type RoomRepository interface {
	Bridger
	ListAllRooms(ctx context.Context) (Rooms, error)
	CreateRoom(ctx context.Context, room *Room) error
	UpdateRoom(ctx context.Context, room Room) error
	ListUserRooms(ctx context.Context, userID string) (Rooms, error)
	GetRoomById(ctx context.Context, roomID string) (Room, error)
	ListRoomParticipants(ctx context.Context, roomID string) ([]RoomParticipant, error)
	DeleteUserRoom(ctx context.Context, roomID, hostID string) error
}
