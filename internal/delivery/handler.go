package delivery

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/encryption"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	redispkg "github.com/elyarsadig/studybud-go/pkg/redis"
	"github.com/elyarsadig/studybud-go/transport"
)

type ApiHandler struct {
	transport.HttpServer
	ctx        context.Context
	logger     logger.Logger
	useCases   map[string]domain.Bridger
	errHandler errorHandler.Handler
	aes        *encryption.AES[string]
	redis      *redispkg.Redis
}

func NewApiHandler(ctx context.Context, aes *encryption.AES[string], redis *redispkg.Redis, errHandler errorHandler.Handler, logger logger.Logger, useCases ...domain.Bridger) (*ApiHandler, error) {
	handler := &ApiHandler{
		useCases:   make(map[string]domain.Bridger),
		ctx:        ctx,
		errHandler: errHandler,
		aes:        aes,
		logger:     logger,
		redis:      redis,
	}

	for _, useCase := range useCases {
		switch useCase.(type) {
		case domain.UserUseCase:
			handler.useCases[configs.USERS_DB_NAME] = useCase
		case domain.TopicUseCase:
			handler.useCases[configs.TOPICS_DB_NAME] = useCase
		case domain.RoomUseCase:
			handler.useCases[configs.ROOMS_DB_NAME] = useCase
		case domain.MessageUseCase:
			handler.useCases[configs.MESSAGES_DB_NAME] = useCase
		}
	}
	return handler, nil
}

func (h *ApiHandler) renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	tmplPaths := []string{
		filepath.Join("web", "main.html"),
		filepath.Join("web", "navbar.html"),
		filepath.Join("web", "activity_component.html"),
		filepath.Join("web", "feed_component.html"),
		filepath.Join("web", "topics_component.html"),
		filepath.Join("web", tmpl),
	}

	tmplParsed, err := template.ParseFiles(tmplPaths...)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}

	err = tmplParsed.ExecuteTemplate(w, "base", data)
	if err != nil {
		h.logger.Error(err.Error())
		return
	}
}

func (h *ApiHandler) ProtectedHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sessionValue, ok := h.extractSessionFromCookie(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		ctx = context.WithValue(ctx, configs.User, sessionValue)
		next(w, r.WithContext(ctx))
	}
}

func (h *ApiHandler) setCookie(w http.ResponseWriter, key string) {
	result, _ := h.aes.Encrypt(key)
	token := base64.URLEncoding.EncodeToString(result)
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  token,
		MaxAge: 3600,
		Secure: true,
	}
	http.SetCookie(w, cookie)
}

func (h *ApiHandler) extractSessionFromCookie(r *http.Request) (domain.SessionValue, bool) {
	ctx := r.Context()
	cookie, err := r.Cookie("session_token")
	if err != nil {
		h.logger.Debug(err.Error())
		return domain.SessionValue{}, false
	}
	token := cookie.Value
	encryptedToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		h.logger.Error(err.Error())
		return domain.SessionValue{}, false
	}
	key, err := h.aes.Decrypt(encryptedToken)
	if err != nil {
		h.logger.Error(err.Error())
		return domain.SessionValue{}, false
	}
	_, jsonData := h.redis.Inspect(ctx, "session", key)
	var sessionValue domain.SessionValue
	err = json.Unmarshal(jsonData, &sessionValue)
	if err != nil {
		h.logger.Error(err.Error())
		return domain.SessionValue{}, false
	}
	return sessionValue, true
}
