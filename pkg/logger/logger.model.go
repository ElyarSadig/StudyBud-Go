package logger

import "github.com/rs/zerolog"

type ZerologLogger struct {
	logger zerolog.Logger
}

type Options struct {
	Development bool // Development add development details of machine
	Debug       bool // Debug show debug devel message
	SkipCaller  int
}
