package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of kubesync",
  Long:  `All software has versions. This is kubesync's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("kubesync bi-direcional kubernetes manifest sync  v0.1 -- HEAD")
  },
}
