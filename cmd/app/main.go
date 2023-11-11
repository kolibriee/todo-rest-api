package main

import "github.com/kostylevdev/todo-rest-api/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
