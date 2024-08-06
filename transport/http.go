package transport

import (
	"context"
	"net/http"
	"time"

	"github.com/elyarsadig/studybud-go/pkg/logger"
	"github.com/go-chi/chi/v5"
)

type HTTPTransporter interface {
	Start()
	Notify() <-chan error
	Shutdown(ctx context.Context) error
	AddHandler(httpMethod HttpMethod, path string, f func(w http.ResponseWriter, r *http.Request))
}

type HttpServer struct {
	server          *http.Server
	router          *chi.Mux
	Address         string
	notify          chan error
	shutDownTimeout time.Duration
	httpAddress     string
}

func NewHTTPServer(httpAddress string, logging logger.Logger) HTTPTransporter {
	newServer := new(HttpServer)
	router := chi.NewRouter()

	httpServer := new(http.Server)
	httpServer.Addr = httpAddress
	httpServer.Handler = PanicRecoverer(router, logging)
	httpServer.WriteTimeout = _defaultWriteTimeout
	httpServer.ReadTimeout = _defaultReadTimeout
	httpServer.ReadHeaderTimeout = _defaultReadHeaderTimeout

	newServer.server = httpServer
	newServer.router = router
	newServer.httpAddress = httpAddress

	return newServer
}

func (s *HttpServer) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *HttpServer) Notify() <-chan error {
	return s.notify
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *HttpServer) AddHandler(httpMethod HttpMethod, path string, f func(w http.ResponseWriter, r *http.Request)) {
	switch httpMethod {
	case POST:
		s.router.Post(path, f)
	case GET:
		s.router.Get(path, f)
	case DELETE:
		s.router.Delete(path, f)
	case PUT:
		s.router.Put(path, f)
	}
}

func PanicRecoverer(h http.Handler, logging logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				logging.Error("panic occured", "panic", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
