package kubernetes

import (
	"github.com/gruntwork-io/terratest/modules/logger"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func ValidateKubernetesConfiguration(t *testing.T, kubeconfigpath string) (*restclient.Config, error) {

	logger.Logf(t, "Validating kubecfg file ", kubeconfigpath)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	if err != nil {
		panic(err.Error())
	}

	return config, err
}
