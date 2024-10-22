package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3.0/agency"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestAgencyAllProjectsLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}

	var agencyId, roleId string
	if agencyId = os.Getenv("OS_AGENCY_ID"); agencyId == "" {
		t.Skip("Agency id required to run this test")
	}
	if roleId = os.Getenv("OS_ROLE_ID"); roleId == "" {
		t.Skip("Role id required to run this test")
	}

	client, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	err = agency.GrantAgencyAllProjects(client, client.DomainID, agencyId, roleId)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		th.AssertNoErr(t, agency.RemoveAgencyAllProjects(client, client.DomainID, agencyId, roleId))
	})

	err = agency.CheckAgencyAllProjects(client, client.DomainID, agencyId, roleId)
	th.AssertNoErr(t, err)

	resp, err := agency.ListAgencyAllProjects(client, client.DomainID, agencyId)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, len(resp.Roles) > 0)
}
