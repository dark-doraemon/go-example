package main

import (
	"project-structure/internal/app"
	"project-structure/internal/config"
)

func main() {

	cfg := config.NewConfig()

	application := app.NewApplication(cfg)

	if err := application.Run(); err != nil {
		panic(err)
	}

}
