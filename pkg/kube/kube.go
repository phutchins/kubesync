package kube

import (
  "fmt"
  "flag"
  errors "errors"
  "path/filepath"
  "github.com/phutchins/kubesync/pkg/config"
  "github.com/phutchins/kubesync/pkg/util"
  kubeErrors "k8s.io/apimachinery/pkg/api/errors"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  corev1 "k8s.io/api/core/v1"
//  v1beta1 "k8s.io/api/extensions/v1beta1"
  appsv1 "k8s.io/api/apps/v1"
  //appsv1 "k8s.io/api/apps/v1"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

func LoadKubeConfig() {
  var kubeConfigFile *string

  if home := util.GetHomeDir(); home != "" {
    kubeConfigFile = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
  } else {
    kubeConfigFile = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
  }
  flag.Parse()

  // use the current context in kubeconfig
  kubeConfig, err := clientcmd.BuildConfigFromFlags("", *kubeConfigFile)
  if err != nil {
    panic(err.Error())
  }

  // create the clientset
  clientset, err = kubernetes.NewForConfig(kubeConfig)
  if err != nil {
    panic(err.Error())
  }

}

func ListNamespaces() (*corev1.NamespaceList) {
  namespaceList, _ := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

  return namespaceList
}

func ListDeployments(namespace string, deployments []string) (appsv1.DeploymentList) {
  //LoadKubeConfig()

  //deploymentLists := make([]appsv1.DeploymentList, 0, 30)

  //namespaces := []string{namespace}

  //for count, namespace := range namespaces {
    deploymentsClient := clientset.AppsV1().Deployments(namespace)

    deploymentList, getErr := deploymentsClient.List(metav1.ListOptions{})
    if getErr != nil {
      panic(fmt.Sprintf("Failed to get latest version of Deployment: %s", getErr))
    }

    //deploymentLists[count] = *deploymentList

    return *deploymentList
  //}
}

func GetPods(ns string, pods []string) (foundPods *corev1.PodList, err error) {
  conf := config.GetConf()
  var errStr string
  err = nil
  foundPods = &corev1.PodList{ListMeta: metav1.ListMeta{ResourceVersion: "0"}}

  if len(pods) == 0 {
    foundPods, err = clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})

    podCount := len(foundPods.Items)
    fmt.Printf("foundPods.Items length is %d\n", podCount)
  } else {
    for index, pod := range pods {
      fmt.Printf("Getting pod %s from index %d\n", pod, index)
      // Examples for error handling:
      // - Use helper functions like e.g. kubeErrors.IsNotFound()
      // - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
      foundPod, err := clientset.CoreV1().Pods(ns).Get(pod, metav1.GetOptions{})
      if kubeErrors.IsNotFound(err) {
        errStr = fmt.Sprintf("Pod %s in namespace %s not found\n", pod, conf.Namespace)
      } else if statusError, isStatus := err.(*kubeErrors.StatusError); isStatus {
        errStr = fmt.Sprintf("Error getting pod %s in namespace %s: %v\n",
          pod, conf.Namespace, statusError.ErrStatus.Message)
      } else if err != nil {
        panic(err.Error())
      } else {
        //errStr = fmt.Sprintf("Found pod %s in namespace %s\n", pods, conf.Namespace)

        foundPods.Items = append(foundPods.Items, *foundPod)

        // Add pod to podList somehow
      }
    }
  }

  // Need to be returning nill for kubeErr if no error is returned
  var kubeErr error
  if errStr != "" {
    kubeErr = errors.New(errStr)
  }

  return foundPods, kubeErr
}
