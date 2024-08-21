package domain

import "context"

type RoomRepository interface {
	Bridger
	ListAllRooms(ctx context.Context) (Rooms, error)
	CreateRoom(ctx context.Context, room *Room) error
	ListUserRooms(ctx context.Context, userID string) (Rooms, error)
}
