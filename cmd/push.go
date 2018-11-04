package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var (
  pushCmd = &cobra.Command{
    Use: "push",
    Short: "Push resources from local to remote",
    RunE: cmdPush,
  }
)

func init() {
  rootCmd.AddCommand(pushCmd)
}

func cmdPush(cmd *cobra.Command, args []string) (err error) {
  fmt.Printf("Got subcmd %v", cmd)

  return
}
