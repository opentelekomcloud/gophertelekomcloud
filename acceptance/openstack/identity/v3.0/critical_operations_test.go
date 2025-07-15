package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3.0/security"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCriticalOperationsLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}
	client, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to GET Operation Protection Policy for domain: %s", client.DomainID)
	opPolicy, err := security.GetOperationProtectionPolicy(client, client.DomainID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, false, *opPolicy.OperationProtection)
	th.AssertEquals(t, "off", opPolicy.AdminCheck)
	th.AssertEquals(t, true, *opPolicy.AllowUser.ManageAccessKey)
	th.AssertEquals(t, true, *opPolicy.AllowUser.ManageEmail)
	th.AssertEquals(t, true, *opPolicy.AllowUser.ManageMobile)
	th.AssertEquals(t, true, *opPolicy.AllowUser.ManagePassword)

	t.Logf("Attempting to Update Operation Protection Policy for domain: %s", client.DomainID)
	opPolicyOpts := security.UpdateProtectionPolicyOpts{
		OperationProtection: pointerto.Bool(true),
		AllowUser: &security.AllowUser{
			ManageAccessKey: pointerto.Bool(false),
			ManageEmail:     pointerto.Bool(false),
			ManageMobile:    pointerto.Bool(false),
			ManagePassword:  pointerto.Bool(false),
		},
	}

	_, err = security.UpdateOperationProtectionPolicy(client, client.DomainID, opPolicyOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Revert Operation Protection Policy to initial state for domain: %s", client.DomainID)
	opPolicyRevertOpts := security.UpdateProtectionPolicyOpts{
		OperationProtection: pointerto.Bool(false),
		AllowUser: &security.AllowUser{
			ManageAccessKey: pointerto.Bool(true),
			ManageEmail:     pointerto.Bool(true),
			ManageMobile:    pointerto.Bool(true),
			ManagePassword:  pointerto.Bool(true),
		},
	}
	_, err = security.UpdateOperationProtectionPolicy(client, client.DomainID, opPolicyRevertOpts)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to GET Operation Protection Policy for domain: %s", client.DomainID)
	opPolicyReverted, err := security.GetOperationProtectionPolicy(client, client.DomainID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, *opPolicyReverted.OperationProtection, *opPolicy.OperationProtection)
}
