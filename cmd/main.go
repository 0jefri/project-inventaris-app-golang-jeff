package main

import "github.com/project-app-inventaris/config"

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}

func main() {

}
