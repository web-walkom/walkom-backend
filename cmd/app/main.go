package main

import (
	"walkom/internal/app"
	"walkom/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	server := new(app.Server)
	if err := server.Run(); err != nil {
		logger.Errorf("Error run server %s", err.Error())
	}
}
