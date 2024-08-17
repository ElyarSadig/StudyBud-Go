package delivery

import "github.com/elyarsadig/studybud-go/internal/domain"

type BaseTemplateData struct {
	Messages []string
	User
}

type User struct {
	IsAuthenticated bool
	AvatarURL       string
	Username        string
}

type Topics struct {
	BaseTemplateData
	domain.Topics
}