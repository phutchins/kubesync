package main

import (
//  "os"
  "fmt"
//  "errors"
//  "path/filepath"
//  "github.com/spf13/cobra"
  "github.com/phutchins/kubesync/cmd"
)

//var (
//  rootCmd = &cobra.Command{
//    Use: "kubesync",
//    Short: "Sync manifests for Kubernetes",
//	  Args: func(cmd *cobra.Command, args []string) error {
//	    if len(args) < 1 {
//	      return errors.New("requires at least one arg")
//	    }
//	    //if myapp.IsValidColor(args[0]) {
//	    //  return nil
//	    //}
//	    //return fmt.Errorf("invalid color specified: %s", args[0])
//      return nil
//	  },
//	  Run: func(cmd *cobra.Command, args []string) {
//	    fmt.Println("Hello, World!")
//	  },
//  }
//
//  defaultConfDir = "$HOME/.kubesync"
//)

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

//  cmd.syncCmd.Flags().String("config",
//    filepath.Join(defaultConfDir, "config"), "path to configuration")
//  cmd.deployCmd.Flags().String("config",
//    filepath.Join(defaultConfDir, "config"), "path to configuration")
  cmd.Execute()
}
