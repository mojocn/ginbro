package parser

import (
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path"
	"text/template"
)

//SaveFlagsToUserConfigFile write flags value to viper config file
func (p *ParseEngine) SaveFlagsToUserConfigFile() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	tplPath := path.Join(goPath, "src", "github.com/dejavuzhou/ginbro", "tpl/cobra.ginbro.yaml")
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return err
	}
	outPath := path.Join(home, "ginbro.yaml")

	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return tmpl.Execute(file, p)

}

//GenerateProjectCode generate go models and handlers file
func (p *ParseEngine) GenerateProjectCode() {
	p.copyStaticSwaggerFolder()
	//make project module directory
	for _, m := range []string{"config", "handlers", "models", "tasks"} {
		p.makeModuleDir(m)
	}
	//copy file
	p.parseTemplateFileChangeImportName()
	//create handlers and models for every table
	for _, tableSchema := range p.Tables {
		if err := tableSchema.generateHandler(); err != nil {
			log.Println(err)
		}
		if err := tableSchema.generateModel("models/model_%s.go"); err != nil {
			log.Println(err)

		}
	}
	p.parseConfigYamlReadmeFiles()
}

//GenerateGormModel just generate GORM model file
func (p *ParseEngine) GenerateGormModel() {
	p.makeModuleDir("")

	//create db helper
	tasks := map[string]string{
		"template/models.db.go.tpl":        "db.go",
		"template/models.db_helper.go.tpl": "db_helper.go",
		"template/models.db_memory.go.tpl": "db_memory.go",
	}
	for kk, vv := range tasks {
		if err := parseTemplate(kk, p.OutPackage, vv, p); err != nil {
			log.Println(err)
		}
	}
	//create models for each tables
	for _, tableSchema := range p.Tables {
		tableSchema.generateModel("model_%s.go")
	}
}
func (p *ParseEngine) parseConfigYamlReadmeFiles() {
	//parse project config.toml
	if err := parseTemplate("tpl/config.toml", p.OutPackage, "config.toml", p); err != nil {
		log.Println(err)
	}
	//parse swagger document yml file
	if err := parseTemplate("tpl/swagger.yaml", p.OutPackage, "swagger/doc.yml", p); err != nil {
		log.Println(err)
	}
	//parse project readme file
	if err := parseTemplate("tpl/readme.md.tpl", p.OutPackage, "readme.md", p); err != nil {
		log.Println(err)
	}
}
func (p *ParseEngine) parseTemplateFileChangeImportName() error {

	tasks := map[string]string{
		"template/tasks.task_example.go.tpl":  "tasks/task_example.go",
		"template/tasks.manager.go.tpl":       "tasks/manager.go",
		"template/tasks.core.go.tpl":          "tasks/core.go",
		"template/tasks.readme.md.tpl":        "tasks/readme.md",
		"template/main.go.tpl":                "main.go",
		"template/config.viper.go.tpl":        "config/viper.go",
		"template/handlers.gin.go.tpl":        "handlers/gin.go",
		"template/handlers.gin_helper.go.tpl": "handlers/gin_helper.go",
		"template/models.db.go.tpl":           "models/db.go",
		"template/models.db_helper.go.tpl":    "models/db_helper.go",
		"template/models.db_memory.go.tpl":    "models/db_memory.go",
		"template/.gitignore.tpl":             ".gitignore",
	}
	for kk, vv := range tasks {
		if err := parseTemplate(kk, p.OutPackage, vv, p); err != nil {
			log.Println(err)
		}
	}
	return nil
}
func (p *ParseEngine) copyStaticSwaggerFolder() {
	tasks := []string{"swagger", "static"}
	for _, vv := range tasks {
		src := path.Join(goPath, "src", "github.com/dejavuzhou/ginbro/boilerplate", vv)
		dst := path.Join(p.OutPath, vv)
		if err := CopyDir(src, dst); err != nil {
			log.Println(err)
		}
	}
}
