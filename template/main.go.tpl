package main

import (
	_ "{{.OutPackage}}/config"
	"{{.OutPackage}}/handlers"
	"{{.OutPackage}}/tasks"
	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
