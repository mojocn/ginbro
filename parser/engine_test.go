package parser

import "testing"

func TestNewParseNgin(t *testing.T) {
	inst := NewParseEngine("root", "password", "127.0.0.1:3306", "dbname", "utf8", "out", "127.0.0.1:9527", "users", "password")
	if inst == nil {
		t.Error("new ParseEngine failed")
	}
}

func TestParseEngine_Close(t *testing.T) {

}

func TestParseEngine_GoFmt(t *testing.T) {

}
func TestParseEngine_BareDatabaseSchema(t *testing.T) {

}

func TestParseEngine_ParseDatabaseSchema(t *testing.T) {

}

func TestParseEngine_GenerateGormModel(t *testing.T) {

}

func TestParseEngine_GenerateProjectCode(t *testing.T) {

}
