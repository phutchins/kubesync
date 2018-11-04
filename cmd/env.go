package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
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
  fmt.Printf("Got subcmd %v", cmd)

  return
}
