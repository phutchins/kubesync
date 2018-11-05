package cmd

import (
  "os"
  "fmt"
  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
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

  if cfgFile != "" {
    fmt.Println("Setting config file to", cfgFile);
    viper.SetConfigFile(cfgFile)
  } else {
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    viper.AddConfigPath(home)
    viper.AddConfigPath(".")
    viper.AddConfigPath("$HOME/")
    //viper.SetConfigFile(".kubesync")
    viper.SetConfigName(".kubesync")
  }
  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("Can't read config:", err)
    os.Exit(1)
  }

  fmt.Println("Config value of context is: ", viper.Get("context"))
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
