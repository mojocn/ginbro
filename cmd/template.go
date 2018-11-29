// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "template",
	Short: "Developer Mode:create template from boilerplate's models and handlers",
	Long:  `Developer Mode:for make your own boilerplate.create template file from boilerplate's models and handlers`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := map[string]string{
			"tasks/core.go":              "tasks.core.go.tpl",
			"tasks/manager.go":           "tasks.manager.go.tpl",
			"tasks/task_example.go":      "tasks.task_example.go.tpl",
			"tasks/readme.md":            "tasks.readme.md.tpl",
			"handlers/gin.go":            "handlers.gin.go.tpl",
			"handlers/gin_helper.go":     "handlers.gin_helper.go.tpl",
			"handlers/middleware_jwt.go": "handlers.middleware_jwt.go.tpl",
			"models/db.go":               "models.db.go.tpl",
			"models/db_helper.go":        "models.db_helper.go.tpl",
			"models/db_memory.go":        "models.db_memory.go.tpl",
			"config/viper.go":            "config.viper.go.tpl",
			"main.go":                    "main.go.tpl",
			".gitignore":                 ".gitignore.tpl",
		}
		for goPath, tplPath := range tasks {
			fileReplace(goPath, tplPath, "github.com/dejavuzhou/ginbro/boilerplate", "{{.OutPackage}}")
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

func fileReplace(goFilepath, tplFilepath, oldString, newString string) error {
	goPath := getGopath()
	goFilepath = path.Join(goPath, "src/github.com/dejavuzhou/ginbro/boilerplate", goFilepath)
	read, err := ioutil.ReadFile(goFilepath)
	if err != nil {
		log.Println(err)
		return err
	}
	newContents := strings.Replace(string(read), oldString, newString, -1)
	tplFilepath = path.Join(goPath, "src/github.com/dejavuzhou/ginbro/template", tplFilepath)
	return ioutil.WriteFile(tplFilepath, []byte(newContents), 777)

}

func getGopath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}
