package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ucs/v1/clusters"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// TestUCSClusterList is read-only and safe to run against any authenticated environment.
func TestUCSClusterList(t *testing.T) {
	client, err := clients.NewUCSV1Client()
	th.AssertNoErr(t, err)

	list, err := clusters.List(client, clusters.ListOpts{})
	th.AssertNoErr(t, err)
	t.Logf("found %d UCS clusters", list.Total)
}

// TestUCSClusterLifecycle registers, queries and deregisters an attached cluster.
// OS_UCS_KUBECONFIG must point to a kubeconfig file of the external cluster to attach.
func TestUCSClusterLifecycle(t *testing.T) {
	kubeconfigPath := clients.EnvOS.GetEnv("UCS_KUBECONFIG")
	if kubeconfigPath == "" {
		t.Skip("OS_UCS_KUBECONFIG (path to a kubeconfig file) is required for the UCS cluster lifecycle test")
	}
	country := clients.EnvOS.GetEnv("UCS_COUNTRY")
	city := clients.EnvOS.GetEnv("UCS_CITY")
	if country == "" || city == "" {
		t.Skip("OS_UCS_COUNTRY and OS_UCS_CITY are required for the UCS cluster lifecycle test")
	}

	kubeconfig, err := os.ReadFile(kubeconfigPath)
	th.AssertNoErr(t, err)

	client, err := clients.NewUCSV1Client()
	th.AssertNoErr(t, err)

	uid, err := clusters.Create(client, clusters.CreateOpts{
		Kind:       "Cluster",
		APIVersion: "v1",
		Metadata: clusters.CreateMetadata{
			Name:        "ucs-acc-attached",
			Annotations: map[string]string{"kubeconfig": string(kubeconfig)},
		},
		Spec: clusters.CreateSpec{
			Category:   "attachedcluster",
			Type:       "privatek8s",
			Provider:   map[string]string{"PRIVATEK8S": "privatek8s"},
			Country:    country,
			City:       city,
			ManageType: "discrete",
		},
	})
	th.AssertNoErr(t, err)
	t.Cleanup(func() {
		th.AssertNoErr(t, clusters.Delete(client, uid))
	})

	got, err := clusters.Get(client, uid)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, uid, got.Metadata.UID)
}
