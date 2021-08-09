package main

import (
	"github.com/yann0917/check-in/cron"
	"github.com/yann0917/check-in/pkg"
)

func main() {
	pkg.Viper()
	cron.Task()
	select {}
}
