package delivery

import "github.com/elyarsadig/studybud-go/internal/domain"

type BaseTemplateData struct {
	Message         string
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

type RoomForm struct {
	Name        string
	Description string
}

type CreateRoomTemplateData struct {
	BaseTemplateData
	TopicList []domain.TopicWithDetails
	Form      RoomForm
}

type UpdateProfileTemplateData struct {
	BaseTemplateData
	Avatar   string
	Name     string
	Username string
	Email    string
	Bio      string
}

type DeleteMessageForm struct {
	BaseTemplateData
	MessageObj string
}

type ActivitiesTemplateData struct {
	BaseTemplateData
	MessageList []domain.Message
}
