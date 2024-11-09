package main

import (
	"github.com/project-app-inventaris/config"
	"github.com/project-app-inventaris/internal/app/delivery"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}

func main() {
	delivery.Server().Run()
}
