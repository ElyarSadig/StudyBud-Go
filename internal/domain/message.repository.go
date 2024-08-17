package domain

import "context"

type MessageRepository interface {
	Bridger
	ListAllMessages(ctx context.Context) (Messages, error)
}
