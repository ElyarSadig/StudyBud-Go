package logger

import (
	"context"
	"errors"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Format int
type LogLevel int

const (
	JSON = iota
	TEXT
)

const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type Logger interface {
	Trace(msg string, args ...any)
	TraceContext(ctx context.Context, msg string, args ...any)
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	Fatal(msg string, args ...any)
	FatalContext(ctx context.Context, msg string, args ...any)
}

func New(outputFormat Format, level LogLevel) (Logger, error) {
	var logger zerolog.Logger

	switch outputFormat {
	case JSON:
		logger = log.Output(os.Stdout)
	case TEXT:
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger = log.Output(consoleWriter)
	default:
		return nil, errors.New("unsoppurted format")
	}

	switch level {
	case TraceLevel:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case DebugLevel:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case InfoLevel:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case WarnLevel:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ErrorLevel:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case FatalLevel:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		return nil, errors.New("unsupported log level")
	}

	logger.Info().
		Str("go_version", runtime.Version()).
		Int("pid", os.Getpid()).
		Str("os", runtime.GOOS).
		Str("os_arch", runtime.GOARCH).
		Msg("Logger initialized")

	return &ZerologLogger{logger: logger}, nil
}

func (l *ZerologLogger) Trace(msg string, args ...any) {
	l.logger.Trace().Fields(args).Msg(msg)
}

func (l *ZerologLogger) TraceContext(ctx context.Context, msg string, args ...any) {
	l.logger.Trace().Ctx(ctx).Fields(args).Msg(msg)
}

func (l *ZerologLogger) Debug(msg string, args ...any) {
	l.logger.Debug().Fields(args).Msg(msg)
}

func (l *ZerologLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.Debug().Ctx(ctx).Fields(args).Msg("")
}

func (l *ZerologLogger) Info(msg string, args ...any) {
	l.logger.Info().Fields(args).Msg(msg)
}

func (l *ZerologLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.Info().Ctx(ctx).Fields(args).Msg(msg)
}

func (l *ZerologLogger) Warn(msg string, args ...any) {
	l.logger.Warn().Fields(args).Msg(msg)
}

func (l *ZerologLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.Warn().Ctx(ctx).Fields(args).Msg(msg)
}

func (l *ZerologLogger) Error(msg string, args ...any) {
	l.logger.Error().Fields(args).Msg(msg)
}

func (l *ZerologLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.Error().Fields(args).Msg(msg)
}

func (l *ZerologLogger) Fatal(msg string, args ...any) {
	l.logger.Fatal().Fields(args).Msg(msg)
}

func (l *ZerologLogger) FatalContext(ctx context.Context, msg string, args ...any) {
	l.logger.Fatal().Ctx(ctx).Fields(args).Msg(msg)
}
