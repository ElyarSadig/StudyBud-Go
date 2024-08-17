package delivery

import (
	"net/http"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
)

func (h *ApiHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "login.html", &BaseTemplateData{})
}

func (h *ApiHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	h.renderTemplate(w, "register.html", &BaseTemplateData{})
}

func (h *ApiHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	data := &BaseTemplateData{}
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
		data.Messages = append(data.Messages, err.Error())
		h.renderTemplate(w, "register.html", data)
		return
	}
	h.setCookie(w, sessionKey)
	
}
