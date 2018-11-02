package main

import (
//  "os"
  "fmt"
  "path/filepath"
  "github.com/spf13/cobra"
)

var (
  rootCmd = &cobra.Command{
    Use: "kubesync",
    Short: "Sync manifests for Kubernetes",
  }

  defaultConfDir = "$HOME/.kubesync"
)

func main() {
  fmt.Sprintf("Starting...")

  // `os.Args` provides access to raw command-line
  // arguments. Note that the first value in this slice
  // is the path to the program, and `os.Args[1:]`
  // holds the arguments to the program.
  //argsWithProg := os.Args
  //argsWithoutProg := os.Args[1:]

  // You can get individual args with normal indexing.
  //arg := os.Args[3]

  //fmt.Println(argsWithProg)
  //fmt.Println(argsWithoutProg)
  //fmt.Println(arg)

  syncCmd.Flags().String("config",
    filepath.Join(defaultConfDir, "config"), "path to configuration")
  deployCmd.Flags().String("config",
    filepath.Join(defaultConfDir, "config"), "path to configuration")
  rootCmd.Execute()
}
