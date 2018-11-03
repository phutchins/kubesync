package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var (
  syncCmd = &cobra.Command{
    Use: "sync",
    Short: "Sync from one direction to the other",
    RunE: cmdSync,
  }
)

func init() {
  rootCmd.AddCommand(syncCmd)
}

func cmdSync(cmd *cobra.Command, args []string) (err error) {
  fmt.Println("Got subcmd %v", cmd)

  return
}
