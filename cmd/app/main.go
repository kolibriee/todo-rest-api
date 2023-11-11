package main

import "github.com/kostylevdev/todo-rest-api/internal/app"

const (
	configName = "config"
	configsDir = "configs"
)

func main() {
	app.Run(configsDir, configName)
}
