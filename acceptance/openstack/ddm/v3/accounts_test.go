package v3

import (
	"testing"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	ddmhelper "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/ddm/v1"
	ddminstancesv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v1/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v3/accounts"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDDMAccountsTestV3(t *testing.T) {
	// CREATE V1 CLIENT
	ddmV1Client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)
	// CREATE V3 CLIENT
	ddmV3Client, err := clients.NewDDMV3Client()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := ddmhelper.CreateDDMInstance(t, ddmV1Client)
	t.Cleanup(func() {
		ddmhelper.DeleteDDMInstance(t, ddmV1Client, ddmInstance.Id)
	})

	// RESET DDM ACCOUNT PASSWORD
	t.Logf("Resetting account password for DDM instance: %s", ddmInstance.Id)
	manageAdminOpts := accounts.ManageAdminPassOpts{
		Name:     "root",
		Password: "acc-test-password1!",
	}
	_, err = accounts.ManageAdminPass(ddmV3Client, ddmInstance.Id, manageAdminOpts)
	th.AssertNoErr(t, err)
	err = golangsdk.WaitFor(600, func() (bool, error) {
		instanceDetails, errP := ddminstancesv1.QueryInstanceDetails(ddmV1Client, ddmInstance.Id)
		th.AssertNoErr(t, errP)
		time.Sleep(5 * time.Second)
		if instanceDetails.Status == "RUNNING" {
			return true, nil
		}
		return false, nil
	})
	th.AssertNoErr(t, err)
}
