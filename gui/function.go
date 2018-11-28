package main

import (
	"github.com/dejavuzhou/ginbro/parser"
	"sync"
)

// Go types that are bound to the UI must be thread-safe, because each binding
// is executed in its own goroutine. In this simple case we may use atomic
// operations, but for more complex cases one should use proper synchronization.
type guiFunction struct {
	sync.Mutex
	result string
}

type args struct {
	MysqlUser     string `json:"mysqlUser"`
	MysqlPassword string `json:"mysqlPassword"`
	MysqlAddr     string `json:"mysqlAddr"`
	MysqlDatabase string `json:"mysqlDatabase"`
	MysqlCharset  string `json:"mysqlCharset"`
	OutPackage    string `json:"outPackage"`
	AppListen     string `json:"appListen"`
	AuthTable     string `json:"authTable"`
	AuthPassword  string `json:"authPassword"`
}

func (c *guiFunction) MysqlGen(arg args) string {
	c.Lock()
	defer c.Unlock()
	ng, err := parser.NewGuiParseEngine(arg.MysqlUser, arg.MysqlPassword, arg.MysqlAddr, arg.MysqlDatabase, arg.MysqlCharset, arg.OutPackage, arg.AppListen, arg.AuthTable, arg.AuthPassword)
	if err != nil {
		return err.Error()
	}
	if err := ng.ParseDatabaseSchema(); err != nil {
		return err.Error()
	}
	ng.GenerateProjectCode()
	ng.GoFmt()
	return "your ginbro project is created at " + ng.OutPath
}
