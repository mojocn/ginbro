package parser

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

type col struct {
	ColumnName string
	//DataType      string
	//ColumnType    string
	ColumnComment string
	ModelProp     string
	ModelType     string
	ModelTag      string
	//ColumnKey     string
	//ColumnType   string
	SwaggerType   string
	SwaggerFormat string
	IsAuthColumn  bool
}

func newCol(cName, dType, cComment, cKey, cType, authColumn string, isAuthTable bool) col {
	cComment = strings.Replace(cComment, `"`, "", -1)
	modelProperty := strcase.ToCamel(strings.ToLower(cName))

	modelType, swgType, swgFormat := transType(dType, cName, cType, cKey)

	//Content   string     `form:"content" json:"content,omitempty" comment:""`
	pt := fmt.Sprintf("%s   %s", modelProperty, modelType)
	sql := ""
	if cType != "" {
		sql = fmt.Sprintf("%s%s", sql, cType)
	}
	if cKey != "" {
		sql = fmt.Sprintf("%s,%s", sql, cKey)
	}
	if dType == "json" {
		cComment = fmt.Sprintf("must a string can unmarsh to an Object. %s", cComment)
	}
	formatTag := `gorm:"column:%s" form:"%s" json:"%s" comment:"%s" sql:"%s"`
	if modelType == "*time.Time" {
		formatTag = `gorm:"column:%s" form:"%s" json:"%s,omitempty" comment:"%s" sql:"%s"`
	}

	tag := fmt.Sprintf(formatTag, cName, cName, cName, cComment, sql)
	isAuthColumn := isAuthTable && cName == authColumn

	modelTag := fmt.Sprintf("%s     `%s`", pt, tag)
	return col{cName, cComment, modelProperty, modelType, modelTag, swgType, swgFormat, isAuthColumn}
}

func transType(dType, cName, cType, cKey string) (string, string, string) {
	modelType := "NoneType"
	swgType, swgFormat := "", ""
	switch dType {
	case "varchar", "longtext", "char", "enum", "set", "mediumtext", "json", "text":
		modelType = "string"
		swgFormat, swgType = "string", "string"
	case "bigint":
		modelType = "int"
		swgFormat, swgType = "int64", "integer"
		if strings.Contains(cType, "unsigned") {
			modelType = "uint64"
			swgFormat, swgType = "int64", "integer"
		}
	case "int", "tinyint", "smallint", "mediumint":
		modelType = "int"
		swgFormat, swgType = "int64", "integer"
		if strings.Contains(cType, "unsigned") {
			modelType = "uint"
			swgFormat, swgType = "int32", "integer"
		}
	case "decimal", "float":
		swgFormat, swgType = "float", "number"
		modelType = "float32"
	case "double":
		swgFormat, swgType = "float", "number"
		modelType = "float64"
	case "blob":
		swgFormat, swgType = "binary", "string"
		modelType = "*[]byte"
	case "time", "datetime", "timestamp":
		swgFormat, swgType = "date-time", "string"
		modelType = "*time.Time"
	}

	if cKey == "PRI" {
		modelType = "uint"
		swgFormat, swgType = "int64", "integer"
	}
	if cName == "ID" || cName == "Id" || cName == "iD" || cName == "id" {
		modelType = "uint"
	}
	return modelType, swgType, swgFormat
}
