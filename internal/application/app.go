package application

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elyarsadig/studybud-go/configs"
	"github.com/elyarsadig/studybud-go/internal/delivery"
	"github.com/elyarsadig/studybud-go/internal/repository"
	"github.com/elyarsadig/studybud-go/internal/usecase"
	confighandler "github.com/elyarsadig/studybud-go/pkg/configHandler"
	"github.com/elyarsadig/studybud-go/pkg/encryption"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	redispkg "github.com/elyarsadig/studybud-go/pkg/redis"
	"github.com/elyarsadig/studybud-go/transport"
	"github.com/hellofresh/health-go/v5"
	"gorm.io/gorm"
)

type Bootstrapper interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

const ApiVersion = "/apis/v1"

type Application struct {
	httpServer        transport.HTTPTransporter
	db                *gorm.DB
	redis             *redispkg.Redis
	logger            logger.Logger
	error             errorHandler.Handler
	serviceConfig     *confighandler.Config[configs.ExtraData]
	healthCheck       *health.Health
	serviceInfo       *configs.ServiceInfo
	sessionExpiration time.Duration
	aes               *encryption.AES[string]
}

func New(
	ctx context.Context,
	httpServer transport.HTTPTransporter,
	errorHandler errorHandler.Handler,
	serviceConfig *confighandler.Config[configs.ExtraData],
	db *gorm.DB,
	redis *redispkg.Redis,
	logger logger.Logger,
	serviceInfo *configs.ServiceInfo,
	aes *encryption.AES[string],
	sessionExpiration time.Duration,
) (Bootstrapper, error) {
	app := new(Application)

	if serviceConfig == nil {
		return nil, errors.New("service config is nil")
	}

	app.aes = aes
	app.db = db
	app.redis = redis
	app.logger = logger
	app.error = errorHandler
	app.httpServer = httpServer
	app.serviceConfig = serviceConfig
	app.serviceInfo = serviceInfo
	app.sessionExpiration = sessionExpiration
	app.healthCheck = healthChecker(serviceInfo.ServiceName, serviceInfo.ServiceVersion, serviceInfo.ServiceCode)

	return app, nil
}

func (a *Application) Run(ctx context.Context) error {
	err := a.registerServiceLayers(ctx)
	if err != nil {
		return err
	}
	a.httpServer.Start()
	a.logger.InfoContext(ctx, "http server has been started",
		"http-address", a.serviceConfig.HttpAddress)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		if err := a.Shutdown(ctx); err != nil {
			panic(err)
		}
		a.logger.WarnContext(ctx, "app/run: signal received", "signal", s.String())
	case err := <-a.httpServer.Notify():
		if err := a.Shutdown(ctx); err != nil {
			panic(err)
		}
		a.logger.WarnContext(ctx, "app/run/httpServer.Notify: ", "error", err)
	}
	return nil
}

func (a *Application) Shutdown(ctx context.Context) error {
	sqlDB, err := a.db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return a.httpServer.Shutdown(ctx)
}

func (a *Application) registerServiceLayers(ctx context.Context) error {
	userRepo := repository.NewUser(a.db, a.error, a.logger)
	topicRepo := repository.NewTopic(a.db, a.error, a.logger)
	roomRepo := repository.NewRoom(a.db, a.error, a.logger)
	messageRepo := repository.NewMessage(a.db, a.error, a.logger)

	userUseCase := usecase.NewUser(a.error, a.sessionExpiration, a.redis, a.logger, userRepo)
	topicUseCase := usecase.NewTopic(a.error, a.logger, topicRepo)
	roomUseCase := usecase.NewRoom(a.error, a.logger, roomRepo, topicRepo)
	messageUseCase := usecase.NewMessage(a.error, a.logger, messageRepo)
	apiHandler, err := delivery.NewApiHandler(ctx, int(a.sessionExpiration.Seconds()), a.aes, a.redis, a.error, a.logger, userUseCase, topicUseCase, roomUseCase, messageUseCase)
	if err != nil {
		return err
	}
	a.registerAPIHandler(apiHandler)

	return nil
}

func (a *Application) registerAPIHandler(apiHandler *delivery.ApiHandler) {
	if a.serviceConfig.ExtraData.HealthCheck {
		a.httpServer.AddHandler("get", "/health", a.healthCheck.HandlerFunc)
	}
	a.httpServer.ServeStaticFiles("web/static")
	a.httpServer.AddHandler("get", "/", apiHandler.HomePage)
	a.httpServer.AddHandler("get", "/logout", apiHandler.Logout)
	a.httpServer.AddHandler("get", "/topics", apiHandler.Topics)
	a.httpServer.AddHandler("get", "/home", apiHandler.HomePage)
	a.httpServer.AddHandler("get", "/room/{id}", apiHandler.RoomPage)
	a.httpServer.AddHandler("post", "/room/{id}", apiHandler.ProtectedHandler(apiHandler.CreateMessage))
	a.httpServer.AddHandler("get", "/activity", apiHandler.ActivitiesPage)
	a.httpServer.AddHandler("get", "/profile/{id}", apiHandler.UserProfilePage)
	a.httpServer.AddHandler("get", "/login", apiHandler.RedirectIfAuthenticated(apiHandler.LoginPage))
	a.httpServer.AddHandler("post", "/login", apiHandler.RedirectIfAuthenticated(apiHandler.LoginUser))
	a.httpServer.AddHandler("get", "/register", apiHandler.RedirectIfAuthenticated(apiHandler.RegisterPage))
	a.httpServer.AddHandler("post", "/register", apiHandler.RedirectIfAuthenticated(apiHandler.RegisterUser))
	a.httpServer.AddHandler("get", "/create-room", apiHandler.ProtectedHandler(apiHandler.CreateRoomPage))
	a.httpServer.AddHandler("post", "/create-room", apiHandler.ProtectedHandler(apiHandler.CreateRoom))
	a.httpServer.AddHandler("get", "/update-room/{id}", apiHandler.ProtectedHandler(apiHandler.UpdateRoomPage))
	a.httpServer.AddHandler("post", "/update-room/{id}", apiHandler.ProtectedHandler(apiHandler.UpdateRoom))
	a.httpServer.AddHandler("get", "/user-update", apiHandler.ProtectedHandler(apiHandler.UpdateProfilePage))
	a.httpServer.AddHandler("post", "/user-update", apiHandler.ProtectedHandler(apiHandler.UpdateProfile))
	a.httpServer.AddHandler("get", "/delete-message/{id}", apiHandler.ProtectedHandler(apiHandler.DeleteMessagePage))
	a.httpServer.AddHandler("post", "/delete-message/{id}", apiHandler.ProtectedHandler(apiHandler.DeleteMessage))
	a.httpServer.AddHandler("get", "/delete-room/{id}", apiHandler.ProtectedHandler(apiHandler.DeleteRoomPage))
	a.httpServer.AddHandler("post", "/delete-room/{id}", apiHandler.ProtectedHandler(apiHandler.DeleteRoom))
}

func healthChecker(name, version, code string) *health.Health {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    fmt.Sprintf("%s - service code: %s", name, code),
		Version: version,
	}))
	return h
}
