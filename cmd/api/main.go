package main

import "tasks-api/cmd/api/modules"

func main() {
	app := modules.NewApp()
	app.Run()
}
