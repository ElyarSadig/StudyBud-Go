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
}

type CreateRoomTemplateData struct {
	BaseTemplateData
	TopicList []domain.TopicWithDetails
	Form      domain.RoomForm
}

type UpdateRoomTemplateData struct {
	BaseTemplateData
	TopicList []domain.TopicWithDetails
	Form      domain.RoomForm
}

type UpdateProfileTemplateData struct {
	BaseTemplateData
	Avatar   string
	Name     string
	Username string
	Email    string
	Bio      string
}

type DeleteForm struct {
	BaseTemplateData
	Obj string
}

type ActivitiesTemplateData struct {
	BaseTemplateData
	MessageList []domain.Message
}

type UserProfileTemplateData struct {
	BaseTemplateData
	TopicList   []domain.TopicWithDetails
	TopicsCount int64
	User        domain.User
	RoomList    []domain.RoomWithDetails
	RoomCount   int64
	MessageList []domain.Message
}

type RoomTemplateData struct {
	BaseTemplateData
	Room         domain.Room
	MessageList  []domain.Message
	Participants []domain.User
}
