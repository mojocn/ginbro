package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	//viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.WithError(err).Error("application configuration'initialization is failed ")
	}
	setLogrus()
}

func setLogrus() {
	lvlString := viper.GetString("app.log_level")
	lvl, err := logrus.ParseLevel(lvlString)
	if err != nil {
		logrus.WithError(err).Fatal("the config file app.log_level only allows debug/info/warn/error/fatal/panic")
	}
	logrus.SetLevel(lvl)
	logrus.RegisterExitHandler(logrusFatalErrorHandler)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	date := time.Now().Format("2006-01-02-app.log")
	f, err := os.OpenFile(date, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(f)
}

func logrusFatalErrorHandler() {
	// gracefully shutdown something...
}
