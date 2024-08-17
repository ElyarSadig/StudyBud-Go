package domain

import "context"

type RoomRepository interface {
	Bridger
	ListAllRooms(ctx context.Context) (Rooms, error)
}
