// Package docs tasks-api.
//
// Documentation of tasks-api.
//
//     	Schemes: http
//     	BasePath: /
//     	Version: 0.0.0
//     	Host: localhost:8080
//     	description: Connects to a local server of tasks-api
//
//     	Consumes:
//     	- application/json
//
//     	Produces:
//     	- application/json
//
//
// swagger:meta
package main

import "tasks-api/cmd/api/modules"

func main() {
	app := modules.NewApp()
	app.Run()
}
