package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yann0917/check-in/cron"
	"github.com/yann0917/check-in/pkg"
)

func main() {
	log.Println("starting...")
	pkg.Viper()
	cron.Task()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("exiting...")
}
