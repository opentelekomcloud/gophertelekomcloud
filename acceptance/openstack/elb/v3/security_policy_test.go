package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/security_policy"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSystemSecurityPolicy(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	systemPolicies, err := security_policy.ListSystemPolicies(client)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, systemPolicies)
}

func TestSecurityPolicyList(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	allPolicies, err := security_policy.List(client, security_policy.ListOpts{})
	th.AssertNoErr(t, err)

	tools.PrintResource(t, allPolicies)
}

func TestSecurityPolicyLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("create-policy-", 3)

	secOpts := security_policy.CreateOpts{
		Name:        policyName,
		Description: "test policy for acceptance testing",
		Protocols: []string{
			"TLSv1",
		},
		Ciphers: []string{
			"AES256-SHA",
		},
	}

	secPolicy, err := security_policy.Create(client, secOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, secPolicy)

	defer deleteSecurityPolicy(t, client, secPolicy.SecurityPolicy.ID)

	updatedName := tools.RandomString("update-policy-", 3)

	updateOpts := security_policy.UpdateOpts{
		Name: updatedName,
	}

	putPolicy, err := security_policy.Update(client, updateOpts, secPolicy.SecurityPolicy.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, putPolicy.SecurityPolicy.Name, updatedName)

	getPolicy, err := security_policy.Get(client, secPolicy.SecurityPolicy.ID)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getPolicy)
	th.AssertEquals(t, getPolicy.SecurityPolicy.ID, secPolicy.SecurityPolicy.ID)
	th.AssertEquals(t, getPolicy.SecurityPolicy.Name, putPolicy.SecurityPolicy.Name)
	th.AssertEquals(t, getPolicy.SecurityPolicy.ProjectId, secPolicy.SecurityPolicy.ProjectId)

	listOpts := security_policy.ListOpts{
		Name: []string{
			updatedName,
		},
	}

	listPolicy, err := security_policy.List(client, listOpts)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, listPolicy)
}

func deleteSecurityPolicy(t *testing.T, client *golangsdk.ServiceClient, secPolicyID string) {
	t.Logf("Attempting to delete ELBv3 Security Policy: %s", secPolicyID)
	err := security_policy.Delete(client, secPolicyID)
	th.AssertNoErr(t, err)
	t.Logf("Deleted ELBv3 security policy: %s", secPolicyID)
}
