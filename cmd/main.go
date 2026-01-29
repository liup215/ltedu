package main

import (
	"edu/conf"
	"edu/lib/logger"
	"edu/server"
	"edu/service"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	logger.InitLogger(conf.Conf.Logger)
	// task.RunTask()

	go func() {
		server.R.Run(fmt.Sprintf(":%s", conf.Conf.Http.Port))
	}()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)

	// Block until a signal is received
	<-quit
	logger.Logger.Info("Shutting down server...")

	service.CloseService()
}
