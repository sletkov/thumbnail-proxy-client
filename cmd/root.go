package cmd

import (
	"fmt"
	"github.com/sletkov/thumbnail-proxy-client/cmd/get"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	RootCmd.AddCommand(get.GetCmd)
}

var RootCmd = &cobra.Command{
	Use:   "thumbnail-proxy-client.exe.exe",
	Short: "thumbnail-proxy-client.exe.exe is a CLI client for thumbnail-proxy service",
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("thumbnail-proxy-client.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("err read config: %v", err)
		return
	}

}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
