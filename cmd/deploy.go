package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

var (
  deployCmd = &cobra.Command{
    Use: "deploy",
    Short: "Deploy a resource",
    RunE: cmdDeploy,
  }
)

func init() {
  rootCmd.AddCommand(deployCmd)
}

func cmdDeploy(cmd *cobra.Command, args []string) (err error) {
  fmt.Printf("Got subcmd %v", cmd)

  return
}
