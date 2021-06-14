package main

import (
	"gitlab.com/dentych/dinner-dash/internal/api"
	"gitlab.com/dentych/dinner-dash/internal/config"
	"gitlab.com/dentych/dinner-dash/internal/database"
	"gitlab.com/dentych/dinner-dash/internal/http"
)

func main() {
	conf := config.FromEnv()
	database.RunMigrations(conf.DbConfig)
	db := database.Init(conf.DbConfig)

	userRepo := database.NewUserRepo(db)
	userApi := api.NewUserApi(userRepo)
	familyApi := api.NewFamilyApi(database.NewFamilyRepo(db), userRepo)
	server := http.NewServer(userApi, familyApi)
	server.SetupAndStart()
}
