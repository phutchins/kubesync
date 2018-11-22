package cmd

import (
	//"fmt"
	"github.com/spf13/cobra"
	//"os"
  "github.com/phutchins/kubesync/pkg/kube"
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

var All bool
var Namespace string

func init() {
	rootCmd.AddCommand(pullCmd)
  rootCmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "Apply to all of this resource")
  rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "default", "The namespace to query")
}

func cmdPull(cmd *cobra.Command, args []string) (err error) {
  var namespaceString string

  // If no args should we pull all?
  if err != nil {
    panic(err.Error())
  }

  //var namespaceString string
  //var namespaces corev1.NamespaceList

  // If no namespace specified, list from all
  //if len(args) == 0 {
    // Add a -a option to list all namespaces
    // Default to default namespace

    // Create Namespace list from this string?
  //  namespaceString = ""
  //} else {
    // Should be type namespacesList ?
    //namespaces = kube.ListNamespaces()
    //namespaceString = args[0]
  //}

  if All == true {
    namespaceString = ""
  } else {
    namespaceString = Namespace
  }

  deploymentList := kube.ListDeployments(namespaceString)

  //for _, deploymentList := range deploymentLists {
    kube.PrintDeployments(deploymentList)
  //}

  /*
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
