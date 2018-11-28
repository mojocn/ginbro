package main

import (
	_ "github.com/dejavuzhou/ginbro/boilerplate/config"
	"github.com/dejavuzhou/ginbro/boilerplate/handlers"
)

func main() {
	defer handlers.Close()
	handlers.ServerRun()
}
