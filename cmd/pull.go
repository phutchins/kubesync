package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var (
  pullCmd = &cobra.Command{
    Use: "pull",
    Short: "Pull resources from remote",
    RunE: cmdPull,
  }
)

func init() {
  rootCmd.AddCommand(pullCmd)
}

func cmdPull(cmd *cobra.Command, args []string) (err error) {
  fmt.Printf("Got subcmd %v", cmd)

  return
}
