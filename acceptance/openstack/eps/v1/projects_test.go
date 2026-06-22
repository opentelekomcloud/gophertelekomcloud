package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/eps/v1/projects"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestEnterpriseProjectsListAndGet(t *testing.T) {
	client, err := clients.NewEPSV1Client()
	th.AssertNoErr(t, err)

	allPages, err := projects.List(client, projects.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)

	allProjects, err := projects.ExtractProjects(allPages)
	th.AssertNoErr(t, err)

	enterpriseProjectID := clients.EnvOS.GetEnv("OS_ENTERPRISE_PROJECT_ID")
	if enterpriseProjectID == "" {
		if len(allProjects) == 0 {
			t.Skip("no enterprise projects are available")
		}
		enterpriseProjectID = allProjects[0].ID
	}

	p, err := projects.Get(client, enterpriseProjectID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, enterpriseProjectID, p.ID)
}
