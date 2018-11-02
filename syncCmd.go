package main

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

func cmdSync(cmd *cobra.Command, args []string) (err error) {
  fmt.Sprintf("hi")

  return
}

func init() {
  rootCmd.AddCommand(syncCmd)
}
