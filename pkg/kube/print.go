package kube

import (
  appsv1 "k8s.io/api/apps/v1"
  corev1 "k8s.io/api/core/v1"
  "fmt"
)

func PrintDeployments(deploymentList appsv1.DeploymentList) {
  for _, d := range deploymentList.Items {
    fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
  }
}

func PrintPods(podList corev1.PodList) {
  for _, p := range podList.Items {
    fmt.Printf(" * %s\n", p.Name)
  }
}
