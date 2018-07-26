package kubernetes

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateKubernetesConfiguration(t *testing.T) {
	t.Parallel()
	kubecfgstring := []byte(`
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlEQ3pDQ0FmT2dBd0lCQWdJUWRzWXp1VTB4UDY1eWNvVE5Nc1VsalRBTkJna3Foa2lHOXcwQkFRc0ZBREF2DQpNUzB3S3dZRFZRUURFeVJsWXpJd1pqazROeTFtTTJJd0xUUmlaamt0T0RFeE55MW1OR1psTVRkaU5qQTBZV1l3DQpIaGNOTVRnd056SXpNakUxTURFM1doY05Nak13TnpJeU1qSTFNREUzV2pBdk1TMHdLd1lEVlFRREV5UmxZekl3DQpaams0TnkxbU0ySXdMVFJpWmprdE9ERXhOeTFtTkdabE1UZGlOakEwWVdZd2dnRWlNQTBHQ1NxR1NJYjNEUUVCDQpBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRQ2V0ZkdDWXRCVXd2QzIxN0xlWWFOOEFURjNMOVN5Nmp5Tm9wMlBraTJHDQpqYlNqRHZhN0ZpT0FjeWw0K3RyM1JBc201SU55dUErVU4xMW1YSnkyaHZ1RDJ2Y1RXalh6K0tvSlRqSDZuTHdjDQpCRmdWKzdEMVJLbERKYUZPVkwwWHhNVkNWK2J4R1hJc3FMRVJVODZXaWI5Qll2OGp6bS9YYTJNL0Z1N3o0UzZHDQpmU2lybWt6eEUvL055WVpKQzZ3a3lMQWNOdGN4WnB3bExHdXFidkttcVM3UTNIeEwyZ1ZvOGFvUGdybXZIbnVHDQp1Z3pOc1g5ckFYbVJxL2UxcHhIVTdKdEQ3SDF6Y2cwNnJDZ3hZNXBQK1BuV2JxcmRrMzUxaktLSjc0di9kV0JpDQoxUmJZcmJxSXJYcEJ3OTdGSDJlM01ycy9mUjZxeFdKeFhPNEtObEtHRmhmUEFnTUJBQUdqSXpBaE1BNEdBMVVkDQpEd0VCL3dRRUF3SUNCREFIRTFHUUUwckpJZUJvMjdqdDBXZ3FMa0NHV21lS0t3QQ0KL0t0bTNZVU5RMlhWdlJUd2NRNHdwMDA5SFdVVXkveGZLN21qbVZ3aVBuZ2ZQQUNjTklUeXY1ajZaZ2lYaW0xVQ0KUGl0WlRWTFZ4Q0VhTkFCdzBxTmdQZkpiVjZDNnQ1Z2VINXNmL2l3WmZsd2I5SllXaXE3SndEc1JiZDdjOC9sYg0KT05QbkQvN041YksrVE1tMWJCRElkc29QMmJobUxKZ1YxOGF1OUpSZUZpbWg2eGFySHNzcmZQclFpZzJ1bmxHWA0KMDgranFqN0pETmhuRjBWN2tueWUNCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0NCg==
    server: https://127.0.0.1
  name: gke_cloud-private-dev_us-central1-b_cde-a95aa8011edb
- cluster:
    certificate-authority: /Users/nobody/.minikube/ca.crt
    server: https://127.0.0.1:8443
  name: minikube
contexts:
- context:
    cluster: gke_cloud-private-dev_us-central1-b_cde-a95aa8011edb
    user: gke_cloud-private-dev_us-central1-b_cde-a95aa8011edb
  name: gke_cloud-private-dev_us-central1-b_cde-a95aa8011edb
- context:
    cluster: minikube
    user: minikube
  name: minikube
current-context: gke_cloud-private-dev_us-central1-b_cde-a95aa8011edb
kind: Config
preferences: {}
users:
- name: gke_cloud-private-dev_us-central1-b_cde-a95aa8011edb
  user:
    auth-provider:
      config:
        access-token: ya29.IaintrealdatahahaIb
        cmd-args: config config-helper --format=json
        cmd-path: /Users/nobody/Downloads/google-cloud-sdk/bin/gcloud
        expiry: 2018-07-23T23:13:52Z
        expiry-key: '{.credential.token_expiry}'
        token-key: '{.credential.access_token}'
      name: gcp
- name: minikube
  user:
    client-certificate: /Users/nobody/.minikube/client.crt
    client-key: /Users/nobody/.minikube/client.key
`)

	tmpfile, err := ioutil.TempFile("", "generic_fake_kubecfg")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(kubecfgstring); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	path, err := filepath.Abs(tmpfile.Name())

	assert.NotPanics(t, func() { ValidateKubernetesConfiguration(t, path) })
}

func TestValidateKubernetesConfigurationWithInvalidConfig(t *testing.T) {
	t.Parallel()

	content := []byte("{}")
	tmpfile, err := ioutil.TempFile("", "generic_invalid_kubecfg")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	path, err := filepath.Abs(tmpfile.Name())
	assert.Panics(t, func() { ValidateKubernetesConfiguration(t, path) }, "The code did not panic")

}
