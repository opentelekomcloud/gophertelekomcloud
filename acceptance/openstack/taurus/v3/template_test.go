package v3

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/instance"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/taurus/v3/template"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTaurusTemplatesList(t *testing.T) {
	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	t.Logf("Attempting to list taurus db configurations")
	getTemplate, err := template.ListConfigurations(client, template.ListConfigurationsOpts{})
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(getTemplate) >= 1, true)
	tools.PrintResource(t, getTemplate)
}

func TestTaurusInstanceTemplatesList(t *testing.T) {
	t.Skip("too long to run in ci")
	vpcID := os.Getenv("OS_VPC_ID")
	subnetID := os.Getenv("OS_NETWORK_ID")
	secGroupID := os.Getenv("OS_SECURITY_GROUP_ID")

	client, err := clients.NewTaurusDBV3Client()
	th.AssertNoErr(t, err)

	createResp := createTaurusInstance(t, client, vpcID, subnetID, secGroupID)

	t.Cleanup(func() {
		t.Logf("Attempting to delete taurus db")
		_, err = instance.Delete(client, createResp.Instance.Id)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to list taurus db configurations")
	getTemplate, err := template.GetInstanceConfigurations(client, template.GetInstanceConfOpts{
		InstanceId: createResp.Instance.Id,
	})
	th.AssertNoErr(t, err)

	th.AssertEquals(t, len(getTemplate) >= 1, true)
	tools.PrintResource(t, getTemplate)
}
