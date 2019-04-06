package util

import (
  "os"
  "fmt"
  "io/ioutil"
  "k8s.io/api/extensions/v1beta1"
  "k8s.io/client-go/kubernetes/scheme"
  "k8s.io/apimachinery/pkg/runtime"
  "encoding/json"
  appsv1 "k8s.io/api/apps/v1"
)

func GetHomeDir() string {
  if h := os.Getenv("HOME"); h != "" {
    return h
  }
  return os.Getenv("USERPROFILE") // windows
}

func LoadJSONFile(filePath string) (obj runtime.Object) {
  decode := scheme.Codecs.UniversalDeserializer().Decode
  jsonFileBytes, err := ioutil.ReadFile(filePath)

  if err != nil {
    fmt.Println(err)
  }

  obj, _, err = decode([]byte(jsonFileBytes), nil, nil)

  if err != nil {
    fmt.Printf("Error decoding JSON file: %#v", err)
  }

  // Maybe we should be pulling in as JSON and merging instead of trying to make a deployment?
  // how do we get kind and verison into the resource?
  //deployment := obj.(*v1beta1.Deployment)
  //fmt.Printf("Type: Deployment Name: %v", deployment.Name)


  return obj
}

func ImportResourceObj(obj runtime.Object) (o interface{}) {
  switch o := obj.(type) {
  case *v1beta1.Deployment:
    fmt.Printf("Type: Deployment Name: %s\n\n", o.ObjectMeta.Name)
    //resource = appsv1.DeploymentSpec{o}
  default:
    fmt.Printf("Unknown type\n")
  }

  //var spec v1beta1.Deployment
  //err = json.Unmarshal(jsonFileBytes, &spec)

  //fmt.Printf("%#v\n", spec)
  return o
}

func JSONEncodeResource(resource appsv1.Deployment) (encodedResource []byte) {
  encodedResource, err := json.MarshalIndent(&resource, "", "\t")
  if err != nil {
    fmt.Println("Err: ", err)
  }

  return encodedResource
}

func EncodeResource(encoding string, resource appsv1.Deployment) (jsonResource []byte, fileExtension string) {
  if encoding == "json" {
    jsonResource = JSONEncodeResource(resource)
    fileExtension = ".json"
  }

  if encoding == "yaml" {
    // How to do one or the other?
  }

  return jsonResource, fileExtension
}

/*
func parseK8sYaml(fileR []byte) []runtime.Object {
  acceptedK8sTypes := regexp.MustCompile(`(Role|ClusterRole|RoleBinding|ClusterRoleBinding|ServiceAccount)`)
  fileAsString := string(fileR[:])
  sepYamlfiles := strings.Split(fileAsString, "---")
  retVal := make([]runtime.Object, 0, len(sepYamlfiles))
  for _, f := range sepYamlfiles {
    if f == "\n" || f == "" {
      // ignore empty cases
      continue
    }

    decode := scheme.Codecs.UniversalDeserializer().Decode
    obj, groupVersionKind, err := decode([]byte(f), nil, nil)

    if err != nil {
      log.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
      continue
    }

    if !acceptedK8sTypes.MatchString(groupVersionKind.Kind) {
      log.Printf("The custom-roles configMap contained K8s object types which are not supported! Skipping object with type: %s", groupVersionKind.Kind)
    } else {
      retVal = append(retVal, obj)
    }

  }
  return retVal
}
*/
