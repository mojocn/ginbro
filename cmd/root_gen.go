package cmd

import (
	"github.com/dejavuzhou/ginbro/parser"
	"github.com/spf13/cobra"
	"log"
)

var appListen, outPackage, authTable, authPassword string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "generate a RESTful APIs app with gin and gorm for gophers",
	Long:    `generate a RESTful APIs app with gin and gorm for gophers`,
	Example: `ginbro gen -u root -p password -a "127.0.0.1:38306" -d dbname -c utf8 --authTable=users --authPassword=pw_column -o=github.com/dejavuzhou/ginbro/out"`,
	Run: func(cmd *cobra.Command, args []string) {
		//SELECT table_name FROM information_schema.tables where table_schema='venom';
		ng := parser.NewParseEngine(mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outPackage, appListen, authTable, authPassword)
		if err := ng.ParseDatabaseSchema(); err != nil {
			log.Println(err)
			cmd.Help()
			return
		}
		ng.GenerateProjectCode()
		ng.GoFmt()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ng := parser.NewParseEngine(mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset, outPackage, appListen, authTable, authPassword)
		err := ng.SaveFlagsToUserConfigFile()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&appListen, "appListen", "l", "127.0.0.1:5555", "app listen Address eg:mojotv.cn, using domain will support gin-TLS")
	genCmd.Flags().StringVarP(&outPackage, "outPackage", "o", "", "output package relative to $GOPATH/src")
	genCmd.Flags().StringVar(&authTable, "authTable", "users", "the MySQL login table")
	genCmd.Flags().StringVar(&authPassword, "authPassword", "password", "password bycrpt column")
	genCmd.MarkFlagRequired("outPackage")
}
