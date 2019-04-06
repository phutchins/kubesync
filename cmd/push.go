package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/phutchins/kubesync/pkg/util"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/util/retry"
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
    localResource := util.ImportResourceObj(obj)

    deploymentName := localResource.ObjectMeta.Name

    retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
      result, getErr := deploymentsClient.get(deploymentName, metav1.GetOptions{})
      if getErr != nil {
        panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
      }

      result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
      result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
      result.Spec = localResource

      _, updateErr := deploymentsClient.Update(result)
      return updateErr
    })
    if retryErr != nil {
      panic(fmt.Errorf("Update failed: %v", retryErr))
    }

    fmt.Println("Updated deployment...")
  }

  return
}
