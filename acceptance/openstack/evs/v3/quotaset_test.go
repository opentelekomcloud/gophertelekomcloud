package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/extensions/quotasets"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumetypes"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQuotasetGet(t *testing.T) {
	clients.RequireAdmin(t)

	client, projectID := getClientAndProject(t)

	quotaSet, err := quotasets.Get(client, projectID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, quotaSet)
}

func TestQuotasetGetDefaults(t *testing.T) {
	clients.RequireAdmin(t)

	client, projectID := getClientAndProject(t)

	quotaSet, err := quotasets.GetDefaults(client, projectID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, quotaSet)
}

func TestQuotasetGetUsage(t *testing.T) {
	clients.RequireAdmin(t)

	client, projectID := getClientAndProject(t)

	quotaSetUsage, err := quotasets.GetUsage(client, projectID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, quotaSetUsage)
}

var UpdateQuotaOpts = quotasets.UpdateOpts{
	Volumes:            golangsdk.IntToPointer(100),
	Snapshots:          golangsdk.IntToPointer(200),
	Gigabytes:          golangsdk.IntToPointer(300),
	PerVolumeGigabytes: golangsdk.IntToPointer(50),
	Backups:            golangsdk.IntToPointer(2),
	BackupGigabytes:    golangsdk.IntToPointer(300),
	Groups:             golangsdk.IntToPointer(350),
	Extra: map[string]interface{}{
		"volumes_foo": golangsdk.IntToPointer(100),
	},
}

var UpdatedQuotas = quotasets.QuotaSet{
	Volumes:            100,
	Snapshots:          200,
	Gigabytes:          300,
	PerVolumeGigabytes: 50,
	Backups:            2,
	BackupGigabytes:    300,
	Groups:             350,
}

var VolumeTypeIsPublic = true
var VolumeTypeCreateOpts = volumetypes.CreateOpts{
	Name:        "foo",
	IsPublic:    &VolumeTypeIsPublic,
	Description: "foo",
	ExtraSpecs:  map[string]string{},
}

func TestQuotasetUpdate(t *testing.T) {
	clients.RequireAdmin(t)

	client, projectID := getClientAndProject(t)

	// save original quotas
	orig, err := quotasets.Get(client, projectID)
	th.AssertNoErr(t, err)

	// create volumeType to test volume type quota
	volumeType, err := volumetypes.Create(client, VolumeTypeCreateOpts)
	th.AssertNoErr(t, err)

	defer func() {
		restore := quotasets.UpdateOpts{}
		FillUpdateOptsFromQuotaSet(*orig, &restore)

		err := volumetypes.Delete(client, volumeType.ID)
		th.AssertNoErr(t, err)

		_, err = quotasets.Update(client, projectID, restore)
		th.AssertNoErr(t, err)

	}()

	// test Update
	resultQuotas, err := quotasets.Update(client, projectID, UpdateQuotaOpts)
	th.AssertNoErr(t, err)

	// We dont know the default quotas, so just check if the quotas are not the
	// same as before
	newQuotas, err := quotasets.Get(client, projectID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, resultQuotas.Volumes, newQuotas.Volumes)
	th.AssertEquals(t, resultQuotas.Extra["volumes_foo"], newQuotas.Extra["volumes_foo"])

	// test that resultQuotas.Extra is populated with the 3 new quota types
	// for the new volumeType foo, don't take into account other volume types
	count := 0
	for k, _ := range resultQuotas.Extra {
		tools.PrintResource(t, k)
		switch k {
		case
			"volumes_foo",
			"snapshots_foo",
			"gigabytes_foo":
			count += 1
		}
	}

	th.AssertEquals(t, count, 3)

	// unpopulate resultQuotas.Extra as it is different per cloud and test
	// rest of the quotaSet
	resultQuotas.Extra = map[string]interface{}(nil)
	th.AssertDeepEquals(t, UpdatedQuotas, *resultQuotas)
}

func TestQuotasetDelete(t *testing.T) {
	clients.RequireAdmin(t)

	client, projectID := getClientAndProject(t)

	// save original quotas
	orig, err := quotasets.Get(client, projectID)
	th.AssertNoErr(t, err)

	defer func() {
		restore := quotasets.UpdateOpts{}
		FillUpdateOptsFromQuotaSet(*orig, &restore)

		_, err = quotasets.Update(client, projectID, restore)
		th.AssertNoErr(t, err)
	}()

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

func FillUpdateOptsFromQuotaSet(src quotasets.QuotaSet, dest *quotasets.UpdateOpts) {
	dest.Volumes = &src.Volumes
	dest.Snapshots = &src.Snapshots
	dest.Gigabytes = &src.Gigabytes
	dest.PerVolumeGigabytes = &src.PerVolumeGigabytes
	dest.Backups = &src.Backups
	dest.BackupGigabytes = &src.BackupGigabytes
	dest.Groups = &src.Groups
}
