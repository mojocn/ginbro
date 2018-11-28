package cmd

import (
	"github.com/dejavuzhou/ginbro/parser"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var bareCmd = &cobra.Command{
	Use:     "bare",
	Short:   "create a bare project",
	Long:    `create a bare project which its mysql flags are not necessary`,
	Example: `ginbro bare -o=github.com/dejavuzhou/ginbro/out5"`,
	Run: func(cmd *cobra.Command, args []string) {
		//SELECT table_name FROM information_schema.tables where table_schema='venom';
		ng := parser.NewParseEngine("mysqlUser", "mysqlPassword", "127.0.0.1:3306", "mysqlDatabase", "utf8", outPackage, "127.0.0.1:9527", "users", "password")
		ng.BareDatabaseSchema()
		ng.GenerateProjectCode()
		ng.GoFmt()
	},
}

func init() {
	rootCmd.AddCommand(bareCmd)
	bareCmd.Flags().StringVarP(&outPackage, "outPackage", "o", "", "output package relative to $GOPATH/src")
	bareCmd.MarkFlagRequired("outPackage")
}
