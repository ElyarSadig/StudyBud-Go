package main

import (
	"log"

	"github.com/elyarsadig/studybud-go/pkg/logger"
)

func main() {
	logger, err := logger.New(logger.JSON, logger.DebugLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug("This is a start!", "start", "end")
}
