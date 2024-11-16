package main

import (
	"fmt"
	_ "test/project/docs"
	"test/project/internal/app"
	"test/project/pkg/config"
)

// @Title Video
// @version 1.0
// @description  API for managing videos.
// @host localhost:2121
// @BasePath /api/test/v1/
func main() {
	configs, err := config.LoadConfiguration()

	if err != nil {
		fmt.Println(err)
	}

	if err = app.InitApp(configs); err != nil {
		fmt.Println("Error in InitApp", err)
	}
}
