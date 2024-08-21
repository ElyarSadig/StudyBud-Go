package domain

import "context"

type MessageUseCase interface {
	Bridger
	ListAllMessages(ctx context.Context) (Messages, error)
	ListUserMessages(ctx context.Context, userID string) (Messages, error)
	ListRoomMessages(ctx context.Context, roomID string) (Messages, error)
	GetUserMessage(ctx context.Context, id string) (Message, error)
	Delete(ctx context.Context, id string) error
}
