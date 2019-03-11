package main
import (
	"log"
	"github.com/urfave/cli"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/golang/glog"
	versionedclient "github.com/f4tq/istio-api-mod/pkg/kube/clientset/versioned"
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v2"
)
func toJsonString(postData interface{}) (string, error) {
	postDataString := new(bytes.Buffer)
	enc := json.NewEncoder(postDataString)

	err := enc.Encode(postData)

	if err != nil {
		return "", err
	}
	return postDataString.String(), nil
}

func fromJsonString(bodyStr string, result interface{}) (err error) {
	var msg json.RawMessage
	body := bytes.NewReader([]byte(bodyStr))
	err = json.NewDecoder(body).Decode(&msg)
	if err == nil {
		switch result.(type) {
		case json.RawMessage:
			tt := result.(*json.RawMessage)
			*tt = msg
		default:
			err = json.Unmarshal(msg, result)

		}
	}
	return err
}

type YamlMsg map[interface{}]interface{}

func fromYamlString(bodyStr string, result interface{}) (err error) {
	return yaml.Unmarshal([]byte(bodyStr), result)
}
func toYamlString(target interface{}) ( string,  error) {
	resultBytes, err := yaml.Marshal(target)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(resultBytes), err
}

func setup(c *cli.Context) (*versionedclient.Clientset,string,error){
	if c.GlobalIsSet("debug") {
		glog.V(8)
	}
	defer glog.Flush()
	kubeconfig := c.GlobalString("kube-config")
	namespace := c.GlobalString("namespace")
	if len(kubeconfig) == 0 || len(namespace) == 0 {
		log.Fatalf("Environment variables KUBECONFIG/--kube-config and NAMESPACE/--namespace need to be set")
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}
	//restConfig.ContentConfig.NegotiatedSerializer
	ic, err := versionedclient.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}

	return ic,namespace,nil

}

