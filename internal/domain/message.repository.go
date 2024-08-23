package domain

import "context"

type MessageRepository interface {
	Bridger
	ListUserMessages(ctx context.Context, userID string) (Messages, error)
	ListRoomMessages(ctx context.Context, roomID string) (Messages, error)
	CreateMessage(ctx context.Context, message *Message) error
	ListAllMessages(ctx context.Context) (Messages, error)
	Get(ctx context.Context, id string) (Message, error)
	Delete(ctx context.Context, id string) error
}
