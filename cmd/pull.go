package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
//	"os"
  "github.com/phutchins/kubesync/pkg/kube"
)

var (
	pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull resources from remote",
		RunE:  cmdPull,
	}
)

var namespace string

func init() {
	rootCmd.AddCommand(pullCmd)
}

func cmdPull(cmd *cobra.Command, args []string) (err error) {
  // If no args should we pull all?
  if len(args) == 0 {
    // GEt a list of all pods
    // Get a single pod
    // gotPods, err := kube.GetPods("*")

    namespace := []string{"automation"}

    list := kube.ListDeployments(namespace)

    fmt.Printf("All deployments:", list)

    return
  }

  pullPod := args[0]

  // Detect what we're pulling or if we're pulling everything

  // Handle wildcards and recursive pull



  fmt.Printf("Pulling pod %s", pullPod)

  pods := []string{pullPod}
  gotPods, err := kube.GetPods(pods)

	if err != nil {
    fmt.Printf("Error: %s", err)
		//panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(gotPods.Items))

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
  /*
	_, err = clientset.CoreV1().Pods(conf.Namespace).Get(pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, conf.Namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, conf.Namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found pod %s in namespace %s\n", pod, conf.Namespace)
	}
  */
  return err
}
