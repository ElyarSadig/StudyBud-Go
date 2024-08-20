package delivery

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/domain"
	"github.com/elyarsadig/studybud-go/pkg/encryption"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	redispkg "github.com/elyarsadig/studybud-go/pkg/redis"
	"github.com/elyarsadig/studybud-go/transport"
	"github.com/google/uuid"
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
		ctx = context.WithValue(ctx, configs.UserCtxKey, sessionValue)
		next(w, r.WithContext(ctx))
	}
}

func (h *ApiHandler) setCookie(w http.ResponseWriter, key string) {
	result, _ := h.aes.Encrypt(key)
	token := base64.URLEncoding.EncodeToString(result)
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  token,
		MaxAge: 5 * 60,
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

func (h *ApiHandler) extractUserProfileUpdateForm(r *http.Request) (domain.UpdateUser, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		return domain.UpdateUser{}, err
	}
	name := r.FormValue("name")
	username := r.FormValue("username")
	bio := r.FormValue("bio")
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		return domain.UpdateUser{}, err
	}
	defer file.Close()
	avatarPath, err := h.saveFileToServer(file, handler)
	if err != nil {
		return domain.UpdateUser{}, err
	}
	user := domain.UpdateUser{
		Name:     name,
		Username: username,
		Bio:      bio,
		Avatar:   avatarPath,
	}
	return user, nil
}

func (h *ApiHandler) saveFileToServer(file multipart.File, handler *multipart.FileHeader) (string, error) {
	uploadDir := "./web/static/uploads"
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filePath := filepath.Join(uploadDir, filename)
	destFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/static/uploads/%s", filename), nil
}
