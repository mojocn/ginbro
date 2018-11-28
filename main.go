// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta

package main

import (
	"github.com/dejavuzhou/ginbro/cmd"
	// preload package for the generated project
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/static"
	_ "github.com/gin-gonic/autotls"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-redis/redis"
	_ "github.com/iancoleman/strcase"
	_ "github.com/jinzhu/gorm"
	_ "github.com/sirupsen/logrus"
	_ "github.com/spf13/viper"
	_ "golang.org/x/crypto/bcrypt"
)

func main() {
	cmd.Execute()
}
