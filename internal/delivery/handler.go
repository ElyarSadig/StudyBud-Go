package delivery

import (
	"context"
	"encoding/base64"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/encryption"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	redispkg "github.com/elyarsadig/studybud-go/pkg/redisPkg"
	"github.com/elyarsadig/studybud-go/transport"
)

type ApiHandler struct {
	transport.HttpServer
	ctx        context.Context
	useCases   map[string]domain.Bridger
	errHandler errorHandler.Handler
	aes        *encryption.AES[string]
	redis      *redispkg.Redis
}

func NewApiHandler(ctx context.Context, aes *encryption.AES[string], redis *redispkg.Redis, errHandler errorHandler.Handler, useCases ...domain.Bridger) (*ApiHandler, error) {
	handler := &ApiHandler{
		useCases:   make(map[string]domain.Bridger),
		ctx:        ctx,
		errHandler: errHandler,
		aes:        aes,
		redis:      redis,
	}

	// for _, useCase := range useCases {
	// 	switch useCase.(type) {
	// 	case domain.UserUseCase:
	// 		handler.useCases[configs.USER_DB_NAME] = useCase
	// 	}
	// }
	return handler, nil
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

func (h *ApiHandler) ProtectedHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		encryptedData, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		sessionValue, err := h.aes.Decrypt(encryptedData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		exists, _ := h.redis.Inspect(ctx, "session", sessionValue)
		if !exists {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		ctx = context.WithValue(ctx, configs.UserName, sessionValue)
		next(w, r.WithContext(ctx))
	}
}
