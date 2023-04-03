package main

import (
	_ "api/docs"
)

// @title Swagger Example API
// @description This is a sample server Petstore server.
// @version 1.0.0
// @host localhost:8010
// @BasePath /
func main() {
	HandleRequest()
}
