package domain

import "context"

type MessageRepository interface {
	Bridger
	ListUserMessages(ctx context.Context, userID string) (Messages, error)
	ListAllMessages(ctx context.Context) (Messages, error)
	Get(ctx context.Context, id string) (Message, error)
	Delete(ctx context.Context, id string) error
}
