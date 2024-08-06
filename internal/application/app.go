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
	confighandler "github.com/elyarsadig/studybud-go/pkg/configHandler"
	"github.com/elyarsadig/studybud-go/pkg/errorHandler"
	"github.com/elyarsadig/studybud-go/pkg/logger"
	"github.com/elyarsadig/studybud-go/transport"
	"github.com/hellofresh/health-go/v5"
	"github.com/redis/go-redis/v9"
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
	redis             *redis.Client
	logger            logger.Logger
	error             errorHandler.Handler
	serviceConfig     *confighandler.Config[configs.ExtraData]
	healthCheck       *health.Health
	serviceInfo       *configs.ServiceInfo
	sessionPrivateKey string
	sessionExpiration time.Duration
}

func New(
	ctx context.Context,
	httpServer transport.HTTPTransporter,
	errorHandler errorHandler.Handler,
	serviceConfig *confighandler.Config[configs.ExtraData],
	db *gorm.DB,
	redis *redis.Client,
	logger logger.Logger,
	serviceInfo *configs.ServiceInfo,
	sessionPrivateKey string,
	sessionExpiration time.Duration,
) (Bootstrapper, error) {
	app := new(Application)

	if serviceConfig == nil {
		return nil, errors.New("service config is nil")
	}

	app.db = db
	app.redis = redis
	app.logger = logger
	app.error = errorHandler
	app.httpServer = httpServer
	app.serviceConfig = serviceConfig
	app.serviceInfo = serviceInfo
	app.sessionPrivateKey = sessionPrivateKey
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
	apiHandler := delivery.NewApiHandler(ctx, a.error, nil)
	a.registerAPIHandler(apiHandler)

	return nil
}

func (a *Application) registerAPIHandler(apiHandler *delivery.ApiHandler) {
	if a.serviceConfig.ExtraData.HealthCheck {
		a.httpServer.AddHandler("get", api("health"), a.healthCheck.HandlerFunc)
	}
}

func api(path string) string {
	out := fmt.Sprintf("%s/%s", ApiVersion, path)
	return out
}

func healthChecker(name, version, code string) *health.Health {

	h, _ := health.New(health.WithComponent(health.Component{
		Name:    fmt.Sprintf("%s - service code: %s", name, code),
		Version: version,
	}))

	return h
}
