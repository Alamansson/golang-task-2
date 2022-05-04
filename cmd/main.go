package main

import (
	golangtask "github.com/alamansson/golang-task-2"
	"github.com/alamansson/golang-task-2/pkg"
	"github.com/sirupsen/logrus"
	"context"
	"os"
	"os/signal"
	"syscall"
)


func main() {
	srv := new(golangtask.Server)
	
	go func() {
		if err := srv.Run( "8000", pkg.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Shutting Down")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
		
	}
}
