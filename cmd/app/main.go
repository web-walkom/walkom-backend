package main

import (
	"github.com/b0shka/walkom-backend/internal/app"
)

func main() {
	server := new(app.Server)
	server.Run()
}
