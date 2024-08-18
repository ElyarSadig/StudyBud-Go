package delivery

import (
	"net/http"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
)

func (h *ApiHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "login.html", BaseTemplateData{})
}

func (h *ApiHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	data := BaseTemplateData{}
	ctx := r.Context()
	email := r.FormValue("email")
	password := r.FormValue("password")
	form := &domain.UserLoginForm{
		Email:    email,
		Password: password,
	}
	useCase := domain.Bridge[domain.UserUseCase](configs.USERS_DB_NAME, h.useCases)
	sessionKey, err := useCase.Login(ctx, form)
	if err != nil {
		data.Message = err.Error()
		h.renderTemplate(w, "login.html", data)
		return
	}
	h.setCookie(w, sessionKey)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func (h *ApiHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "register.html", BaseTemplateData{})
}

func (h *ApiHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	data := BaseTemplateData{}
	ctx := r.Context()
	name := r.FormValue("name")
	userName := r.FormValue("username")
	email := r.FormValue("email")
	password1 := r.FormValue("password1")
	password2 := r.FormValue("password2")
	form := &domain.UserRegisterForm{
		Name:      name,
		Username:  userName,
		Email:     email,
		Password1: password1,
		Password2: password2,
	}
	useCase := domain.Bridge[domain.UserUseCase](configs.USERS_DB_NAME, h.useCases)
	sessionKey, err := useCase.RegisterUser(ctx, form)
	if err != nil {
		data.Message = err.Error()
		h.renderTemplate(w, "register.html", data)
		return
	}
	h.setCookie(w, sessionKey)
}

func (h *ApiHandler) Topics(w http.ResponseWriter, r *http.Request) {
	data := BaseTemplateData{}
	ctx := r.Context()
	useCase := domain.Bridge[domain.TopicUseCase](configs.TOPICS_DB_NAME, h.useCases)
	queryParams := r.URL.Query()
	name := queryParams.Get("q")
	var topics domain.Topics
	var err error
	if len(name) == 0 {
		topics, err = useCase.ListAllTopics(ctx)
		if err != nil {
			data.Message = err.Error()
			h.renderTemplate(w, "topics.html", data)
			return
		}
	} else {
		topics, err = useCase.SearchTopicByName(ctx, name)
		if err != nil {
			data.Message = err.Error()
			h.renderTemplate(w, "topics.html", data)
			return
		}
	}
	tmplData := Topics{
		BaseTemplateData: data,
		Topics:           topics,
	}
	h.renderTemplate(w, "topics.html", tmplData)
}

func (h *ApiHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	data := HomeTemplateData{}
	data.RequestUser.Username = h.extractUserNameFromCookie(r)
	ctx := r.Context()
	topicUseCase := domain.Bridge[domain.TopicUseCase](configs.TOPICS_DB_NAME, h.useCases)
	roomUseCase := domain.Bridge[domain.RoomUseCase](configs.ROOMS_DB_NAME, h.useCases)
	messageUseCase := domain.Bridge[domain.MessageUseCase](configs.MESSAGES_DB_NAME, h.useCases)
	topics, err := topicUseCase.ListAllTopics(ctx)
	if err != nil {
		data.Message = err.Error()
		h.renderTemplate(w, "home.html", data)
		return
	}
	data.TopicList = topics.List
	data.TopicsCount = topics.Count
	rooms, err := roomUseCase.ListAllRooms(ctx)
	if err != nil {
		data.Message = err.Error()
		h.renderTemplate(w, "home.html", data)
		return
	}
	data.RoomCount = rooms.Count
	data.RoomList = rooms.List
	messages, err := messageUseCase.ListAllMessages(ctx)
	if err != nil {
		data.Message = err.Error()
		h.renderTemplate(w, "home.html", data)
		return
	}
	data.MessageList = messages.MessageList
	h.renderTemplate(w, "home.html", data)
}
