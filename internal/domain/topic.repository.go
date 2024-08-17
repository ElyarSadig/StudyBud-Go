package domain

import "context"

type TopicRepository interface {
	Bridger
	ListAllTopics(ctx context.Context) (Topics, error)
	SearchTopicByName(ctx context.Context, name string) (Topics, error)
}
