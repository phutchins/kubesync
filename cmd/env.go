package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/phutchins/kubesync/pkg/config"
)

var (
  envCmd = &cobra.Command{
    Use: "env",
    Short: "Set environment params",
    RunE: cmdEnv,
  }
)

func init() {
  rootCmd.AddCommand(envCmd)
}

func cmdEnv(cmd *cobra.Command, args []string) (err error) {
  conf := config.GetConf()

  fmt.Println("Current environment values are:")

  fmt.Printf("current-context: %s\n", conf.CurrentContext)
  fmt.Printf("namespace:       %s\n", conf.Namespace)

  return
}
