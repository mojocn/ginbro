package parser

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"log"
)

type schema struct {
	ResourceName         string
	HandlerName          string
	TableName            string
	ModelName            string
	Columns              []col
	ProjectPath          string
	OutPackage           string
	IsAuthTable          bool
	PasswordPropertyName string
	PasswordColumnName   string
}

func newTable(tableName, projectPath, OutPackage, authTable, authPassword string) *schema {
	mName := strcase.ToCamel(tableName)
	mName = inflection.Singular(mName)
	rName := strcase.ToKebab(mName)
	hName := strcase.ToLowerCamel(mName)
	isAuthTable := tableName == authTable
	passwordPropName := strcase.ToCamel(authPassword)
	return &schema{ModelName: mName, TableName: tableName, ResourceName: rName, HandlerName: hName, ProjectPath: projectPath, OutPackage: OutPackage,
		IsAuthTable: isAuthTable, PasswordColumnName: authPassword, PasswordPropertyName: passwordPropName}
}

func (t *schema) generateModel(nameFormat string) error {
	if nameFormat == "" {
		nameFormat = "models/model_%s.go"
	}
	fileName := fmt.Sprintf(nameFormat, t.TableName)
	if t.IsAuthTable {
		if err := parseTemplate("tpl/models.jwt.go.tpl", t.OutPackage, "models/jwt.go", t); err != nil {
			log.Println(err)
		}
	}
	return parseTemplate("tpl/models.template.tpl", t.OutPackage, fileName, t)
}

func (t *schema) generateHandler() error {
	fileName := fmt.Sprintf("handlers/handler_%s.go", t.TableName)
	if t.IsAuthTable {
		err := parseTemplate("tpl/handlers.handler_auth.go.tpl", t.OutPackage, "handlers/handler_auth.go", t)
		if err != nil {
			log.Println(err)
		}
		err = parseTemplate("template/handlers.middleware_jwt.go.tpl", t.OutPackage, "handlers/middleware_jwt.go", t)
		if err != nil {
			log.Println(err)
		}
	}

	return parseTemplate("tpl/handlers.template.tpl", t.OutPackage, fileName, t)
}
