package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	ddmhelper "github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack/ddm/v1"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ddm/v3/accounts"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDDMAccountsTestV3(t *testing.T) {
	//CREATE V1 CLIENT
	ddmv1client, err := clients.NewDDMV1Client()
	th.AssertNoErr(t, err)
	// CREATE V3 ClIENT
	ddmv3client, err := clients.NewDDMV3Client()
	th.AssertNoErr(t, err)

	// CREATE DDM INSTANCE
	ddmInstance := ddmhelper.CreateDDMInstance(t, ddmv1client)
	t.Cleanup(func() {
		ddmhelper.DeleteDDMInstance(t, ddmv1client, ddmInstance.Id)
	})

	// RESET DDM ACCOUNT PASSWORD
	t.Logf("Resetting account password for DDM instance: %s", ddmInstance.Id)
	manageAdminOpts := accounts.ManageAdminPassOpts{
		Name:     "root",
		Password: "acc-test-password1!",
	}
	_, err = accounts.ManageAdminPass(ddmv3client, ddmInstance.Id, manageAdminOpts)
	th.AssertNoErr(t, err)

}
