package delivery

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/go-chi/chi/v5"
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

func (h *ApiHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	encryptedData, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	key, err := h.aes.Decrypt(encryptedData)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	err = h.redis.Remove(ctx, "session", key)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
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
	sessionValue, ok := h.extractSessionFromCookie(r)
	if ok {
		data = BaseTemplateData{
			AvatarURL:       sessionValue.Avatar,
			Username:        sessionValue.Username,
			IsAuthenticated: true,
		}
	}
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
	sessionValue, ok := h.extractSessionFromCookie(r)
	data := HomeTemplateData{
		BaseTemplateData: BaseTemplateData{
			Username:        sessionValue.Username,
			IsAuthenticated: ok,
			AvatarURL:       sessionValue.Avatar,
		},
	}
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

func (h *ApiHandler) CreateRoomPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sessionValue := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	data := CreateRoomTemplateData{
		BaseTemplateData: BaseTemplateData{
			IsAuthenticated: true,
			Username:        sessionValue.Username,
			AvatarURL:       sessionValue.Avatar,
		},
	}
	useCase := domain.Bridge[domain.TopicUseCase](configs.TOPICS_DB_NAME, h.useCases)
	topics, err := useCase.ListAllTopics(ctx)
	if err != nil {
		data.Message = err.Error()
		h.renderTemplate(w, "room_form.html", data)
		return
	}
	data.TopicList = topics.List
	h.renderTemplate(w, "room_form.html", data)
}

func (h *ApiHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// sessionValue := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	// data := BaseTemplateData{
	// 	AvatarURL:       sessionValue.Avatar,
	// 	Username:        sessionValue.Username,
	// 	IsAuthenticated: true,
	// }
	// useCase := domain.Bridge[domain.RoomUseCase](configs.ROOMS_DB_NAME, h.useCases)
	// err := useCase.CreateRoom(ctx, RoomForm, sessionValue.ID)
	// if err != nil {
	// 	data.Message = err.Error()
	// 	h.renderTemplate(w, "room_form.html", data)
	// 	return
	// }
	// http.Redirect(w, r, "/home", http.StatusFound)
}

func (h *ApiHandler) UpdateProfilePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sessionValue := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	useCase := domain.Bridge[domain.UserUseCase](configs.USERS_DB_NAME, h.useCases)
	user, err := useCase.GetUserByEmail(ctx, sessionValue.Email)
	if err != nil {
		data := BaseTemplateData{}
		data.Message = err.Error()
		h.renderTemplate(w, "update_user.html", data)
		return
	}
	data := UpdateProfileTemplateData{
		BaseTemplateData: BaseTemplateData{
			IsAuthenticated: true,
			Username:        sessionValue.Username,
			AvatarURL:       sessionValue.Avatar,
		},
		Avatar:   user.Avatar,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Bio:      user.Bio,
	}
	h.renderTemplate(w, "update_user.html", data)
}

func (h *ApiHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	updateUser, err := h.extractUserProfileUpdateForm(r)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	useCase := domain.Bridge[domain.UserUseCase](configs.USERS_DB_NAME, h.useCases)
	sessionKey, err := useCase.UpdateInfo(ctx, &updateUser)
	if err != nil {
		data := BaseTemplateData{Message: err.Error()}
		h.renderTemplate(w, "update_user.html", data)
		return
	}
	h.setCookie(w, sessionKey)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func (h *ApiHandler) DeleteMessagePage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request := ctx.Value(configs.UserCtxKey).(domain.SessionValue)
	baseData := BaseTemplateData{
		IsAuthenticated: true,
		AvatarURL:       request.Avatar,
		Username:        request.Username,
	}
	useCase := domain.Bridge[domain.MessageUseCase](configs.MESSAGES_DB_NAME, h.useCases)
	id := chi.URLParam(r, "id")
	message, err := useCase.GetUserMessage(ctx, id)
	if err != nil {
		h.renderTemplate(w, "not_found.html", baseData)
		return
	}
	data := DeleteMessageForm{
		BaseTemplateData: baseData,
		MessageObj:       message.Body,
	}
	h.renderTemplate(w, "delete_message.html", data)
}

func (h *ApiHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	useCase := domain.Bridge[domain.MessageUseCase](configs.MESSAGES_DB_NAME, h.useCases)
	err := useCase.Delete(ctx, id)
	if err != nil {
		h.renderTemplate(w, "delete_message.html", BaseTemplateData{Message: err.Error()})
		return
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}

func (h *ApiHandler) ActivitiesPage(w http.ResponseWriter, r *http.Request) {
	sessionValue, ok := h.extractSessionFromCookie(r)
	baseData := BaseTemplateData{
		AvatarURL:       sessionValue.Avatar,
		Username:        sessionValue.Username,
		IsAuthenticated: ok,
	}
	ctx := r.Context()
	useCase := domain.Bridge[domain.MessageUseCase](configs.MESSAGES_DB_NAME, h.useCases)
	messages, err := useCase.ListAllMessages(ctx)
	if err != nil {
		baseData.Message = err.Error()
		h.renderTemplate(w, "activity.html", baseData)
		return
	}
	data := ActivitiesTemplateData{
		BaseTemplateData: baseData,
		MessageList:      messages.MessageList,
	}
	h.renderTemplate(w, "activity.html", data)
}
