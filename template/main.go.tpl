package main

import (
	_ "{{.OutPackage}}/config"
	"{{.OutPackage}}/handlers"
)

func main() {
	defer handlers.Close()
	handlers.ServerRun()
}
