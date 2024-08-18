package delivery

import "github.com/elyarsadig/studybud-go/internal/domain"

type BaseTemplateData struct {
	Messages        []string
	IsAuthenticated bool
	AvatarURL       string
	Username        string
}

type Topics struct {
	BaseTemplateData
	domain.Topics
}

type HomeTemplateData struct {
	BaseTemplateData
	TopicList   []domain.TopicWithDetails
	TopicsCount int64
	RoomList    []domain.RoomWithDetails
	RoomCount   int64
	MessageList []domain.Message
	RequestUser domain.User
}
