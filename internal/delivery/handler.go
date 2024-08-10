package delivery

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/transport"
)

type ApiHandler struct {
	transport.HttpServer
	ctx        context.Context
	useCases   map[string]domain.Bridger
	errHandler errorHandler.Handler
}

func NewApiHandler(ctx context.Context, errHandler errorHandler.Handler, useCases ...domain.Bridger) *ApiHandler {
	handler := &ApiHandler{
		useCases:   make(map[string]domain.Bridger),
		ctx:        ctx,
		errHandler: errHandler,
	}

	// for _, useCase := range useCases {
	// 	switch useCase.(type) {
	// 	case domain.UserUseCase:
	// 		handler.useCases[configs.USER_DB_NAME] = useCase
	// 	}
	// }
	return handler
}

func (h *ApiHandler) renderTemplate(w http.ResponseWriter, tmpl string, data *BaseTemplateData) error {
	tmplPaths := []string{
		filepath.Join("web", "main.html"),
		filepath.Join("web", "navbar.html"),
		filepath.Join("web", tmpl),
	}

	tmplParsed, err := template.ParseFiles(tmplPaths...)
	if err != nil {
		return err
	}

	err = tmplParsed.ExecuteTemplate(w, "base", data)
	if err != nil {
		return err
	}
	return nil
}

func (h *ApiHandler) Login(w http.ResponseWriter, r *http.Request) {
	data := &BaseTemplateData{}
	h.renderTemplate(w, "login.html", data)
}

func (h *ApiHandler) Register(w http.ResponseWriter, r *http.Request) {
	data := &BaseTemplateData{}
	h.renderTemplate(w, "register.html", data)
}
