package delivery

import (
	"context"

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
