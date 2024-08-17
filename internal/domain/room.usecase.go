package domain

import "context"

type RoomUseCase interface {
	Bridger
	ListAllRooms(ctx context.Context) (Rooms, error)
}
