package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
  "github.com/phutchins/kubesync/pkg/kube"
  appsv1 "k8s.io/api/apps/v1"
  //corev1 "k8s.io/api/core/v1"

  // Use this for decoding yaml and jason files
  //"k8s.io/apimachinery/pkg/util/yaml"
)

// pull without a destination will list resources that will be pulled
var (
	pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull resources from remote",
		RunE:  cmdPull,
	}
)

var (
  pullDeploymentCmd = &cobra.Command{
    Use:   "deployment",
    Short: "Deployment resource",
    RunE:  cmdPullDeployments,
  }
)

var All bool
var Namespace string

func init() {
	rootCmd.AddCommand(pullCmd)
  rootCmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "Apply to all of this resource")
  rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "default", "The namespace to query")

  pullCmd.AddCommand(pullDeploymentCmd)
}

func cmdPull(cmd *cobra.Command, args []string) (err error) {
  fmt.Println("Will pull all resources...")

  return err
}

func cmdPullDeployments(cmd *cobra.Command, args []string) (err error) {
  var namespaceString string

  if err != nil {
    panic(err.Error())
  }

  if All == true {
    namespaceString = ""
  } else {
    namespaceString = Namespace
  }

  deploymentStrings := args

  err, deploymentList := PullDeployments(namespaceString, deploymentStrings)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  kube.PrintDeployments(deploymentList)

  return err
}

// Make this a sub command of pull which will pull deployments
func PullDeployments(namespaceString string, deploymentStrings []string) (err error, deploymentList appsv1.DeploymentList) {

  deploymentList = kube.ListDeployments(namespaceString, deploymentStrings)

  return err, deploymentList
}

func PullPods (ns *string, podStrings *[]string) (err error) {
  // Handle wildcards and recursive pull

  gotPods, err := kube.GetPods(*podStrings)

	if err != nil {
    fmt.Printf("Error: %s", err)
		panic(err.Error())
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


