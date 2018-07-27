package kubernetes

import (
	"github.com/gruntwork-io/terratest/modules/logger"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func LoadKubernetesConfiguration(t *testing.T, kubeconfigpath string) (*restclient.Config, error) {

	logger.Logf(t, "Validating kubecfg file ", kubeconfigpath)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	if err != nil {
		panic(err.Error())
	}

	return config, err
}

func ListPods(t *testing.T, kubeconfigpath string) (*v1.PodList, error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)

	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		logger.Logf(t, "There are %d pods in the cluster\n", len(pods.Items))
		return pods, nil

	}

}
