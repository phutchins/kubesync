package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kubesync [OPTIONS] [COMMANDS]",
	Short: "Kubesync is a bi-directional sync utility for Kubernetes",
	Long: `Kubesync provides convenience, safety and extra functionality around
                managing your Kubernetes resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubesync)")
}

func initConfig() {
  viper.SetConfigType("yaml")

  // Set ENV prefix and load ENV variables
  viper.SetEnvPrefix("ks")
  viper.BindEnv("env")

  if cfgFile != "" {
    fmt.Println("Setting config file to", cfgFile);
    viper.SetConfigFile(cfgFile)
  } else {
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Println("Home: ", home)

    // Should find a way to load config file from configPaths without an extension
    //viper.SetConfigName(".kubesync")
    //viper.AddConfigPath("/Users/philip/")
    //viper.AddConfigPath(home)
    //viper.AddConfigPath(".")
    //viper.AddConfigPath("$HOME/")
    viper.SetConfigFile("/Users/philip/.kubesync")
  }

  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("Can't read config:", err)
    os.Exit(1)
  }

  fmt.Println("Config value of current-context is: ", viper.Get("current-context"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
