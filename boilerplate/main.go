package main

import (
	_ "github.com/dejavuzhou/ginbro/boilerplate/config"
	"github.com/dejavuzhou/ginbro/boilerplate/handlers"
	"github.com/dejavuzhou/ginbro/boilerplate/tasks"
	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
