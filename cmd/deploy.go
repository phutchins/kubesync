package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var (
  deployCmd = &cobra.Command{
    Use: "deploy",
    Short: "deploy from one direction to the other",
    RunE: cmdDeploy,
  }
)

func cmdDeploy(cmd *cobra.Command, args []string) (err error) {
  fmt.Sprintf("hi")

  return
}

func init() {
  rootCmd.AddCommand(deployCmd)
}
