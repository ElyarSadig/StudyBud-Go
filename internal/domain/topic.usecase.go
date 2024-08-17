package domain

import "context"

type TopicUseCase interface {
	Bridger
	ListAllTopics(ctx context.Context) (Topics, error)
	SearchTopicByName(ctx context.Context, name string) (Topics, error)
}
