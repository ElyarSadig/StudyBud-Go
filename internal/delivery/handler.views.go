package delivery

import (
	"net/http"
)

func (h *ApiHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	data := &BaseTemplateData{}
	h.renderTemplate(w, "login.html", data)
}

func (h *ApiHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	data := &BaseTemplateData{}
	h.renderTemplate(w, "register.html", data)
}

func (h *ApiHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// data := &BaseTemplateData{}
	// name := r.FormValue("name")
	// userName := r.FormValue("username")
	// email := r.FormValue("email")
	// password1 := r.FormValue("password1")
	// password2 := r.FormValue("password2")
	// if password1 != password2 {
	// 	data.Messages = append(data.Messages, "passwords do not match!")
	// 	h.renderTemplate(w, "register.html", data)
	// 	return
	// }
}
