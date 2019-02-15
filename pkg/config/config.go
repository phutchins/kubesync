package config

import (
  "fmt"
  "os"
  "github.com/spf13/viper"
  //homedir "github.com/mitchellh/go-homedir"
)

type conf struct {
  CurrentContext string
  Namespace string
  RootPath string
}

func main() {

}

// Should take cfgFile param here
func InitConfig () {
  viper.SetConfigType("yaml")

  // Set ENV prefix and load ENV variables
  viper.SetEnvPrefix("ks")
  viper.BindEnv("env")

/*
  if cfgFile != "" {
    fmt.Println("Setting config file to", cfgFile);
    viper.SetConfigFile(cfgFile)
  } else {
/*
/*
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
*/

    // Should find a way to load config file from configPaths without an extension
    //viper.SetConfigName(".kubesync")
    //viper.AddConfigPath("/Users/philip/")
    //viper.AddConfigPath(home)
    //viper.AddConfigPath(".")
    //viper.AddConfigPath("$HOME/")
    viper.SetConfigFile("/Users/philip/.kubesync")
//  }

  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("Can't read config:", err)
    os.Exit(1)
  }
}

func GetConf() conf {
  return conf{
    CurrentContext: viper.Get("current-context").(string),
    Namespace: viper.Get("namespace").(string),
    RootPath: viper.Get("rootPath").(string)}
}

