package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/phutchins/kubesync/pkg/util"
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
  pushNames := args

  for _, pushName := range pushNames {
    fmt.Printf("Pushing %v\n", pushName)

    obj := util.LoadJSONFile(pushName)
    util.ImportResourceObj(obj)
  }

  return
}
