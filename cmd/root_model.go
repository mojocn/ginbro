package cmd

import (
	"github.com/dejavuzhou/ginbro/parser"
	"github.com/spf13/cobra"
	"log"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:     "model",
	Short:   "only generate GORM models of all MySQL tables",
	Long:    `generate GORM models of MySQL tables.`,
	Example: `ginbro model -u root -p password -a 127.0.0.1:3306 -d venom -c utf8  -o=github.com/dejavuzhou/ginbro/out_model`,
	Run: func(cmd *cobra.Command, args []string) {
		ng := parser.NewParseEngine(mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outPackage, "", "", "")
		if err := ng.ParseDatabaseSchema(); err != nil {
			log.Println(err)
			cmd.Help()
			return
		}
		ng.GenerateGormModel()
		ng.GoFmt()
		log.Println("you have to edit the db.go file of viper initialize")
		log.Println("doc https://github.com/spf13/viper ")
		log.Println("demo https://github.com/dejavuzhou/ginbro/blob/master/boilerplate/config/viper.go")
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
	modelCmd.Flags().StringVarP(&outPackage, "outPackage", "o", "", "eg: models,the models will be created at $GOPATH/src/models")
	modelCmd.MarkFlagRequired("outPackage")

}
