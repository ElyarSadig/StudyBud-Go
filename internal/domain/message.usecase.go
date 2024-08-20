package domain

import "context"

type MessageUseCase interface {
	Bridger
	ListAllMessages(ctx context.Context) (Messages, error)
	GetUserMessage(ctx context.Context, id string) (Message, error)
}
