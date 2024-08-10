package delivery

type BaseTemplateData struct {
	Messages []string
	User
}

type User struct {
	IsAuthenticated bool
	AvatarURL       string
	Username        string
}
