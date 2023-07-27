package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
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

func TestSecurityPolicyLifecycle(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("create-policy-", 3)

	secPolicy := createSecurityPolicy(t, client, policyName)
	tools.PrintResource(t, secPolicy)

	t.Cleanup(func() { deleteSecurityPolicy(t, client, secPolicy.Id) })

	updatedName := tools.RandomString("update-policy-", 3)

	updateOpts := security_policy.UpdateOpts{
		Name: updatedName,
	}

	putPolicy, err := security_policy.Update(client, updateOpts, secPolicy.Id)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, putPolicy.Name, updatedName)

	getPolicy, err := security_policy.Get(client, secPolicy.Id)
	th.AssertNoErr(t, err)

	tools.PrintResource(t, getPolicy)
	th.AssertEquals(t, getPolicy.Id, secPolicy.Id)
	th.AssertEquals(t, getPolicy.Name, putPolicy.Name)
	th.AssertEquals(t, getPolicy.ProjectId, secPolicy.ProjectId)

	allPages, err := security_policy.List(client, security_policy.ListOpts{
		Name: []string{
			updatedName,
		},
	}).AllPages()
	th.AssertNoErr(t, err)

	allPolicies, err := security_policy.ExtractSecurities(allPages)
	th.AssertNoErr(t, err)

	for _, pool := range allPolicies {
		tools.PrintResource(t, pool)
	}
}

func TestPolicyAssignment(t *testing.T) {
	client, err := clients.NewElbV3Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("create-policy-", 3)

	loadbalancerID := createLoadBalancer(t, client)
	t.Cleanup(func() { deleteLoadbalancer(t, client, loadbalancerID) })

	certificateID := createCertificate(t, client)
	t.Cleanup(func() { deleteCertificate(t, client, certificateID) })

	t.Run("AssignSecurityPolicyListenerCreation", func(t *testing.T) {
		secPolicyID := createSecurityPolicy(t, client, policyName).Id
		t.Cleanup(func() { deleteSecurityPolicy(t, client, secPolicyID) })

		listenerName := tools.RandomString("create-listener-", 3)

		createOpts := listeners.CreateOpts{
			DefaultTlsContainerRef: certificateID,
			Description:            "some interesting description",
			LoadbalancerID:         loadbalancerID,
			Name:                   listenerName,
			Protocol:               "HTTPS",
			ProtocolPort:           443,
			SecurityPolicy:         secPolicyID,
		}

		listener, err := listeners.Create(client, createOpts)
		t.Cleanup(func() {
			t.Logf("Attempting to delete ELBv3 Listener: %s", listener.ID)
			err := listeners.Delete(client, listener.ID)
			th.AssertNoErr(t, err)
			t.Logf("Deleted ELBv3 Listener: %s", listener.ID)
		})
		th.AssertNoErr(t, err)
		th.AssertEquals(t, listener.SecurityPolicy, secPolicyID)
	})

	t.Run("AssignSecurityPolicyListenerUpdate", func(t *testing.T) {
		secPolicyUpdatedID := createSecurityPolicy(t, client, policyName).Id
		t.Cleanup(func() { deleteSecurityPolicy(t, client, secPolicyUpdatedID) })
		listenerName := tools.RandomString("create-listener-", 3)

		createOpts := listeners.CreateOpts{
			DefaultTlsContainerRef: certificateID,
			Description:            "some interesting description",
			LoadbalancerID:         loadbalancerID,
			Name:                   listenerName,
			Protocol:               "HTTPS",
			ProtocolPort:           443,
		}

		listener, err := listeners.Create(client, createOpts)
		th.AssertNoErr(t, err)
		t.Cleanup(func() {
			t.Logf("Attempting to delete ELBv3 Listener: %s", listener.ID)
			err := listeners.Delete(client, listener.ID)
			th.AssertNoErr(t, err)
			t.Logf("Deleted ELBv3 Listener: %s", listener.ID)
		})

		updateOpts := listeners.UpdateOpts{
			SecurityPolicy: secPolicyUpdatedID,
		}

		_, _ = listeners.Update(client, listener.ID, updateOpts)

		updatedListener, err := listeners.Get(client, listener.ID)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, updatedListener.SecurityPolicy, secPolicyUpdatedID)
	})
}

func deleteSecurityPolicy(t *testing.T, client *golangsdk.ServiceClient, secPolicyID string) {
	t.Logf("Attempting to delete ELBv3 Security Policy: %s", secPolicyID)
	err := security_policy.Delete(client, secPolicyID)
	th.AssertNoErr(t, err)
	t.Logf("Deleted ELBv3 security policy: %s", secPolicyID)
}

func createSecurityPolicy(t *testing.T, client *golangsdk.ServiceClient, policyName string) *security_policy.SecurityPolicy {
	t.Logf("Attempting to create ELBv3 security policy")
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
	t.Logf("Created  ELBv3 security policy: %s", secPolicy.Id)

	return secPolicy
}
