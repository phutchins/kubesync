package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
  "github.com/phutchins/kubesync/pkg/kube"
  "encoding/json"
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

var (
  pullDeploymentsCmd = &cobra.Command{
    Use:   "deployments",
    Short: "Deployments resource",
    RunE:  cmdPullDeployments,
  }
)

var (
  pullPodCmd = &cobra.Command{
    Use:  "pod",
    Short: "Pod resource",
    RunE: cmdPullPods,
  }
)

var (
  pullPodsCmd = &cobra.Command{
    Use:  "pods",
    Short: "Pods resource",
    RunE: cmdPullPods,
  }
)

var All bool
var Format string
var Output string
var Destination string
var Namespace string

func init() {
	rootCmd.AddCommand(pullCmd)
  rootCmd.PersistentFlags().BoolVarP(&All, "all", "a", false, "Query all namespaces")
  rootCmd.PersistentFlags().StringVarP(&Format, "format", "f", "json", "Set format of the output")
  // Output is either file or stdout
  rootCmd.PersistentFlags().StringVarP(&Output, "output", "o", "stdout", "Set destination for output")
  // Instead of output and destination, if destination is set, output is to file, otherwise it is to stdout

  // Output can be determined by the options given
  // - if there is a ./ or a path we can assume output is to file and that location
  // - if no output location given, output should be stdout
  // File output destination
  rootCmd.PersistentFlags().StringVarP(&Destination, "destination", "d", "./", "Set file location")
  rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "default", "The namespace to query")

  pullCmd.AddCommand(pullDeploymentCmd)
  pullCmd.AddCommand(pullDeploymentsCmd)
  pullCmd.AddCommand(pullPodCmd)
  pullCmd.AddCommand(pullPodsCmd)
}

func cmdPull(cmd *cobra.Command, args []string) (err error) {
  fmt.Println("Will pull all resources...")

  return err
}

func checkArgsForPath(args []string) (path string) {
  if len(args) > 1 {
    path = args[1]
  }

  return path
}

func cmdPullDeployments(cmd *cobra.Command, args []string) (err error) {
  var namespaceString string

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

  if len(deploymentList.Items) == 0 {
    fmt.Println("No deployments found");

    return err
  }

  kube.PrintDeployments(deploymentList)

  filePath := checkArgsForPath(args)

  for _, deployment := range deploymentList.Items {
    //err := json.NewEncoder(mDeployment).Encode(deployment)
    mDeployment, _ := json.MarshalIndent(&deployment, "", "\t")

    // Check to see if we want to display or save to disk

    // If we display just print
    fmt.Println("deployment: ", string(mDeployment))

    if Output != "stdout" {
      fmt.Println("Got destination arg")
    }

    if filePath != "" {
      fmt.Println("filepath exists: ", filePath)
      err := writeToFile(filePath, mDeployment)
      if err != nil {
        fmt.Println("Error writing to file: %s", err)
      }
    }


    // If we save to disk
      // Convert each deployment object to json
      // Determine where in the file structure this file should go
      // Look for existing file and load if it exists
        // If it exists load it
          // Diff the pulled file and loaded file
      // Save json to file
  }

  // return err
  return err
}

func writeToFile(filePath string, b []byte) (err error) {
  file, err := os.Create(filePath)
  file.Write(b)
  defer file.Close()

  return err
}

func cmdPullPods(cmd *cobra.Command, args []string) (err error) {
  var namespaceString string

  if All == true {
    namespaceString = ""
  } else {
    namespaceString = Namespace
  }

  err = PullPods(&namespaceString, &args)

  return err
}

// Make this a sub command of pull which will pull deployments
func PullDeployments(namespaceString string, deploymentStrings []string) (err error, deploymentList appsv1.DeploymentList) {

  err, deploymentList = kube.ListDeployments(namespaceString, deploymentStrings)

  if err != nil {
    fmt.Printf("Error getting deployment list: %s\n", err)
    os.Exit(1)
  }

  return err, deploymentList
}

func PullPods (ns *string, podStrings *[]string) (err error) {
  // Handle wildcards and recursive pull
  gotPods, err := kube.GetPods(*ns, *podStrings)

	if err != nil {
    fmt.Printf("Error: %s\n", err)
    os.Exit(1)
	}

	fmt.Printf("There are %d pods in the cluster\n", len(gotPods.Items))

  kube.PrintPods(*gotPods)

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


