package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/quotasets"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQuotasetGet(t *testing.T) {
	client, projectID := getClientAndProject(t)

	quotaSet, err := quotasets.Get(client, projectID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, quotaSet)
}

func TestQuotasetGetDefaults(t *testing.T) {
	client, projectID := getClientAndProject(t)

	quotaSet, err := quotasets.GetDefaults(client, projectID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, quotaSet)
}

func TestQuotasetGetUsage(t *testing.T) {
	client, projectID := getClientAndProject(t)

	quotaSetUsage, err := quotasets.GetUsage(client, projectID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, quotaSetUsage)
}

func TestQuotasetDelete(t *testing.T) {
	client, projectID := getClientAndProject(t)

	// save original quotas
	_, err := quotasets.Get(client, projectID)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		restore := quotasets.UpdateOpts{}
		_, err = quotasets.Update(client, projectID, restore)
		th.AssertNoErr(t, err)
	})

	// Obtain environment default quotaset values to validate deletion.
	defaultQuotaSet, err := quotasets.GetDefaults(client, projectID)
	th.AssertNoErr(t, err)

	// Test Delete
	err = quotasets.Delete(client, projectID)
	th.AssertNoErr(t, err)

	newQuotas, err := quotasets.Get(client, projectID)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, newQuotas.Volumes, defaultQuotaSet.Volumes)
}

// getClientAndProject reduces boilerplate by returning a new blockstorage v3
// ServiceClient and a project ID obtained from the OS_PROJECT_NAME envvar.
func getClientAndProject(t *testing.T) (*golangsdk.ServiceClient, string) {
	client, err := clients.NewBlockStorageV3Client()
	th.AssertNoErr(t, err)

	projectID := os.Getenv("OS_PROJECT_NAME")
	th.AssertNoErr(t, err)
	return client, projectID
}
