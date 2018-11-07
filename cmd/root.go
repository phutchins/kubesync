package cmd

import (
	"fmt"
	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
  "github.com/phutchins/kubesync/pkg/config"
  "github.com/phutchins/kubesync/pkg/kube"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kubesync [OPTIONS] [COMMANDS]",
	Short: "Kubesync is a bi-directional sync utility for Kubernetes",
	Long: `Kubesync provides convenience, safety and extra functionality around
                managing your Kubernetes resources.`,
	Run: func(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
      cmd.Help()
      os.Exit(0)
    }
	},
}

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubesync)")
  // Should be passing config file name override flag here
	cobra.OnInitialize(config.InitConfig)
  kube.LoadKubeConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
