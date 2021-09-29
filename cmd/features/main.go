package main

import (
	"features/config"
	"features/internal/app"
	"features/util"
)

func main() {
	conf, err := config.Load(".")
	util.PanicIfError(err, "failed to load config")
	log := config.SetUpLogger(conf)
	log.Info("Initializing server")

	application := app.NewApp(conf, log)
	application.Run()
}
