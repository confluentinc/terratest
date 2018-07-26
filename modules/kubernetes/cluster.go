package kubernetes

import (	"k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"testing"
	"github.com/stretchr/testify/assert"
	restclient "k8s.io/client-go/rest"
)


func ValidateKubernetesConfiguration(t *testing.T, kubeconfigpath string) (*restclient.Config, error) {
	t.Parallel()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	if err != nil {
		panic(err.Error())
	}


	assert.NoError(t, err)
	assert.NotEmpty(t, config)
	return config, err
}

