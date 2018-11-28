package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfgFile, mysqlUser, mysqlPassword, mysqlAddr, mysqlDatabase, mysqlCharset string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ginbro",
	Short: "Ginbro is a scaffold tool for Gin and GORM, one command to generate a mighty RESTful APIs App",
	Long: `
fastest way to generate a RESTful APIs application with MySQL in Go
support JWT Authorization Bearer Auth and JWT middleware
support brute-force-login firewall
build in swift golang-memory cache
generate GORM model from MySQL database schema
powered with Swagger document and SwaggerUI
capable of serve VueJs app's static files
configurable CORS middleware
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("create called")
	//
	//	},

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println(mysqlDatabase)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/ginbro.yaml)")
	rootCmd.PersistentFlags().StringVarP(&mysqlUser, "mysqlUser", "u", "root", "MySQL user name")
	rootCmd.PersistentFlags().StringVarP(&mysqlPassword, "mysqlPassword", "p", "password", "MySQL password")
	rootCmd.PersistentFlags().StringVarP(&mysqlAddr, "mysqlAddr", "a", "127.0.0.1:3306", "MySQL host:port")
	rootCmd.PersistentFlags().StringVarP(&mysqlDatabase, "mysqlDatabase", "d", "", "MySQL database name")
	rootCmd.PersistentFlags().StringVarP(&mysqlCharset, "mysqlCharset", "c", "utf8", "MySQL charset")
	viper.BindPFlag("mysqlDatabase", rootCmd.PersistentFlags().Lookup("mysqlDatabase"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name "ginbro" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("ginbro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
