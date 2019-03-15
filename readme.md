# [Converting a MySQL database'schema to a RESTful golang APIs app in the fastest way](https://github.com/dejavuzhou/ginbro)
[![Build Status](https://travis-ci.org/dejavuzhou/ginbro.svg?branch=master)](https://travis-ci.org/dejavuzhou/ginbro) 
[![GoDoc](http://godoc.org/github.com/dejavuzhou/ginbro?status.svg)](http://godoc.org/github.com/dejavuzhou/ginbro) 
[![Go Report Card](https://goreportcard.com/badge/github.com/dejavuzhou/ginbro)](https://goreportcard.com/report/github.com/dejavuzhou/ginbro)
![stability-stable](https://img.shields.io/badge/stability-stable-brightgreen.svg)
[![codebeat badge](https://codebeat.co/badges/650029a5-fcea-4416-925e-277e2f178e96)](https://codebeat.co/projects/github-com-dejavuzhou-ginbro-master)
[![codecov](https://codecov.io/gh/dejavuzhou/ginbro/branch/master/graph/badge.svg)](https://codecov.io/gh/dejavuzhou/ginbro)

Ginbro is a scaffold tool for Gin-Gorm-MySQL which you just need to input one command to create a mighty RESTful APIs App.
## Warning
### - this Repo's code has transfered to [Felix](https://github.com/dejavuzhou/felix), please visit[dejavuzhou/felix](https://github.com/dejavuzhou/felix)
### - 代码已经转移到[dejavuzhou/felix](https://github.com/dejavuzhou/felix),请移步至[dejavuzhou/felix](https://github.com/dejavuzhou/felix)

## Demo and Translated Document
- [中文文档](readme_zh.md)            
- [Video-Demo-Youtube](https://www.youtube.com/watch?v=TvWQhNKfmCo&feature=youtu.be)
- [Video-Demo-Bilibili](https://www.bilibili.com/video/av36804258/)

## Feature
- [fastest way to generate a RESTful APIs application with MySQL in Go](/boilerplate)
- support [JWT Authorization Bearer](boilerplate/handlers/middleware_jwt.go) [Auth](boilerplate/handlers/handler_auth.go) and [JWT middleware](boilerplate/models/jwt.go)
- [support brute-force-login firewall](boilerplate/models/model_users.go)
- [build in swift golang-memory cache](https://github.com/dejavuzhou/ginbro/blob/master/boilerplate/models/db_memory.go)
- [generate GORM model from MySQL database schema](boilerplate/models)
- [powered with Swagger document and SwaggerUI](boilerplate/swagger)
- [capable of serve VueJs app's static files](boilerplate/static)
- [configurable CORS middleware](boilerplate/handlers/gin_helper.go)
- [user friendly configuration](tpl/config.toml)
- [golang GUI app](gui)
- [fully build-in cron task support](boilerplate/tasks)
- [travis CI/CD](https://travis-ci.org/dejavuzhou/ginbro)
    
## Ginbro Installation
you can install it by `go get` command：
```shell
go get github.com/dejavuzhou/ginbro
```
the Ginbro executable binary will locate in $GOPATH/bin
[check GOBIN is in your environment PATH](https://stackoverflow.com/questions/25216765/gobin-not-set-cannot-run-go-install)

## Usage

### 1. `ginbro gen` generate a new Gin+Gorm+MySQL RESTful APIs Application with JWT middleware and auth
example 

`ginbro gen -u root -p Password -a "127.0.0.1:3306" -d databasename -o "github.com/user/awesome" -c utf8 --authTable=users --authPassword=password`
```bash

$ ginbro gen -h
generate a RESTful APIs app with gin and gorm for gophers

Usage:
  ginbro gen [flags]

Examples:
ginbro gen -u root -p password -a "127.0.0.1:38306" -d dbname -c utf8 --authTable=users --authPassword=pw_column -o=github.com/dejavuzhou/ginbro/out"

Flags:
  -l, --appListen string      app listen Address eg:mojotv.cn, using domain will support gin-TLS (default "127.0.0.1:5555")
      --authPassword string   password bycrpt column (default "password")
      --authTable string      the MySQL login table (default "users")
  -h, --help                  help for gen
  -o, --outPackage string     output package relative to $GOPATH/src

Global Flags:
      --config string          config file (default is $HOME/ginbro.yaml)
  -a, --mysqlAddr string       MySQL host:port (default "127.0.0.1:3306")
  -c, --mysqlCharset string    MySQL charset (default "utf8")
  -d, --mysqlDatabase string   MySQL database name
  -p, --mysqlPassword string   MySQL password (default "password")
  -u, --mysqlUser string       MySQL user name (default "root")
```
#### the generated project directory [ginbro DEMO-code-repository](https://github.com/dejavuzhou/ginbro-son)

### 2. `ginbro bare` generate a bare project with one resource which you have to edit the `config.toml` which is easy for you to customize
```bash
$ ginbro bare -h
create a bare project which its mysql flags are not necessary

Usage:
  ginbro bare [flags]

Examples:
ginbro bare -o=github.com/dejavuzhou/ginbro/out5"

Flags:
  -h, --help                help for bare
  -o, --outPackage string   output package relative to $GOPATH/src
```
### 3. `ginbro model` generate GORM models of tables in a MySQL database
```bash
$ genbro model -h
generate GORM models of MySQL tables.

Usage:
  ginbro model [flags]

Examples:
ginbro model -u root -p password -a 127.0.0.1:3306 -d venom -c utf8  -o=github.com/dejavuzhou/ginbro/out_model

Flags:
  -h, --help                help for model
  -o, --outPackage string   eg: models,the models will be created at $GOPATH/src/models

Global Flags:
      --config string          config file (default is $HOME/ginbro.yaml)
  -a, --mysqlAddr string       MySQL host:port (default "127.0.0.1:3306")
  -c, --mysqlCharset string    MySQL charset (default "utf8")
  -d, --mysqlDatabase string   MySQL database name
  -p, --mysqlPassword string   MySQL password (default "password")
  -u, --mysqlUser string       MySQL user name (default "root")
```
## [GUI](/gui)

## [Boilerplate Project](/boilerplate)

## environment
- my development environment
    - Windows 10 pro 64
    - go version go1.11.1 windows/amd64
    - mysql version <= 5.7

## go packages
```shell
go get github.com/gin-contrib/cors
go get github.com/gin-contrib/static
go get github.com/gin-gonic/autotls
go get github.com/gin-gonic/gin
go get github.com/sirupsen/logrus
go get github.com/spf13/viper
go get github.com/spf13/cobra
go get github.com/go-redis/redis
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
go get github.com/dgrijalva/jwt-go
```
## How to fix `go get golang.org/x/crypto/bcrypt` and `go get golang.org/x/crypto/text` error
```bash
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/crypto
git clone https://github.com/golang/text
```
retry the command`go get github.com/dejavuzhou/ginbro`

## Info
- resource table'schema which has no "ID","id","Id'" or "iD" will not generate model or route.
- the column which type is json value must be a string which is able to decode into a JSON, when resource is called POST or PATCH from the swaggerUI.
## Thanks
- [swagger Specification](https://swagger.io/specification/)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [GORM](http://gorm.io/)
- [viper](https://github.com/spf13/viper)
- [cobra](https://github.com/spf13/cobra#getting-started)
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- [base64captcha](https://github.com/mojocn/base64Captcha)
## Please feedback your [`issue`](https://github.com/dejavuzhou/ginbro/issues) with database schema file
