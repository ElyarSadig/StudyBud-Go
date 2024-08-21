package domain

import "context"

type RoomUseCase interface {
	Bridger
	ListAllRooms(ctx context.Context) (Rooms, error)
	CreateRoom(ctx context.Context, form RoomForm) error
}
