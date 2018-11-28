package parser

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/exec"
	"path"
)

//ParseEngine parser engine
type ParseEngine struct {
	db            *sql.DB
	database      string
	OutPath       string
	OutPackage    string
	MysqlAddr     string
	MysqlUser     string
	MysqlPassword string
	MysqlDatabase string
	MysqlCharset  string
	AppSecret     string
	AppListen     string
	Tables        []*schema
	AuthTable     string
	AuthPassword  string
}

//NewParseEngine create a parser engine
func NewGuiParseEngine(mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outDir, appListen, authTable, authColumn string) (*ParseEngine, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	projectPath := path.Join(goPath, "src", outDir)
	return &ParseEngine{
		db,
		mysqlDatabase,
		projectPath,
		outDir,
		mysqlAddr,
		mysqlUser,
		mysqlPassword,
		mysqlDatabase,
		mysqlCharset,
		randomString(32),
		appListen,
		[]*schema{},
		authTable,
		authColumn,
	}, nil
}

func NewParseEngine(mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outDir, appListen, authTable, authColumn string) *ParseEngine {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		log.Println(err)
	}
	projectPath := path.Join(goPath, "src", outDir)
	return &ParseEngine{
		db,
		mysqlDatabase,
		projectPath,
		outDir,
		mysqlAddr,
		mysqlUser,
		mysqlPassword,
		mysqlDatabase,
		mysqlCharset,
		randomString(32),
		appListen,
		[]*schema{},
		authTable,
		authColumn,
	}
}

//Close the mysql database
func (p *ParseEngine) Close() {
	p.db.Close()
}

func (p *ParseEngine) makeModuleDir(module string) {
	modulePath := path.Join(p.OutPath, module)
	if err := os.MkdirAll(modulePath, 0777); err != nil {
		log.Fatal(err)
	}
}

//GoFmt run gofmt for the generated project
func (p *ParseEngine) GoFmt() {
	log.Println("running go fmt for the new project")

	runCmd("go", "fmt", p.OutPackage+"/...")
	//mainPath := path.Join(p.OutPath,"main.go")
	//runCmd("go","run",mainPath)
	log.Printf("cd %s", p.OutPath)
	log.Println("test your project")

}

func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("%s", out)
	}
}
