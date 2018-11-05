package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var (
  configureCmd = &cobra.Command{
    Use: "configure",
    Short: "Configure kubesync",
    RunE: cmdConfigure,
  }
)

func init() {
  rootCmd.AddCommand(configureCmd)
}

func cmdConfigure(cmd *cobra.Command, args []string) (err error) {
  fmt.Printf("Got subcmd %v", cmd)

  return
}
