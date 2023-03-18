package main

import (
	"github.com/b0shka/walkom-backend/internal/app"
	"github.com/b0shka/walkom-backend/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	server := new(app.Server)
	if err := server.Run(); err != nil {
		logger.Errorf("Error run server %s", err.Error())
	}
}
