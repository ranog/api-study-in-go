package main

import (
	"github.com/ranog/api-study-in-go/config/logger"
	"log/slog"
)

func main() {
	logger.InitLogger()

	user := logger.User{
		Name:     "John",
		Age:      30,
		Password: "123456",
	}

	slog.Info("Starting API")
	slog.Info("Creating user", "user", user.LogUser())
}
