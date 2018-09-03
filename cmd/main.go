package main

import (
	"context"
	"gitlab.com/dentych/dinner-dash/internal/config"
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/http"
)

func main() {
	conf := config.FromEnv()
	database.RunMigrations(conf.DbConfig)
	database.Init(conf.DbConfig)

	http.Setup(context.Background())
}
