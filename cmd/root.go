package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// var create string
// var overlays string
var rootCmd = &cobra.Command{
	Use:   "pass-auth",
	Short: "This program can create a password and check overlays.",
	Long:  `This program can create a password with your specifications and check keyboard overlays for password.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("Can't start root cmd, error : %s", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

}

func initConfig() {
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//TODO create and add command "overlays"
