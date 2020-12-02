package vpnaas

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/extensions/vpnaas/ikepolicies"
)

func TestIkePolicyList(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a NetworkingV2 client: %s", err)
	}

	listOpts := ikepolicies.ListOpts{}
	allPages, err := ikepolicies.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to fetch Ike Policy pages: %s", err)
	}
	ikePolicies, err := ikepolicies.ExtractPolicies(allPages)
	if err != nil {
		t.Fatalf("Unable to extract Ike Policy pages: %s", err)
	}
	for _, policy := range ikePolicies {
		tools.PrintResource(t, policy)
	}
}

func TestIkePolicyLifeCycle(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a NetworkingV2 client: %s", err)
	}

	// Create Ike Policy
	ikePolicy, err := createIkePolicy(t, client)
	if err != nil {
		t.Fatalf("Unable to create Ike Policy: %s", err)
	}
	defer deleteIkePolicy(t, client, ikePolicy.ID)

	tools.PrintResource(t, ikePolicy)

	err = updateIkePolicy(t, client, ikePolicy.ID)
	if err != nil {
		t.Fatalf("Unable to update Ike Policy: %s", err)
	}
	tools.PrintResource(t, ikePolicy)

	newIkePolicy, err := ikepolicies.Get(client, ikePolicy.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to get Ike Policy: %s", err)
	}
	tools.PrintResource(t, newIkePolicy)
}

func createIkePolicy(t *testing.T, client *golangsdk.ServiceClient) (*ikepolicies.Policy, error) {
	policyName := tools.RandomString("create-ike-", 8)

	createIkePolicyOpts := ikepolicies.CreateOpts{
		Description:           "some ike policy description",
		Name:                  policyName,
		AuthAlgorithm:         ikepolicies.AuthAlgorithm("md5"),
		EncryptionAlgorithm:   ikepolicies.EncryptionAlgorithm("3des"),
		PFS:                   ikepolicies.PFS("group1"),
		Phase1NegotiationMode: ikepolicies.Phase1NegotiationMode("main"),
		IKEVersion:            ikepolicies.IKEVersion("v1"),
		Lifetime: &ikepolicies.LifetimeCreateOpts{
			Units: ikepolicies.Unit("seconds"),
			Value: 1800,
		},
	}
	ikePolicy, err := ikepolicies.Create(client, createIkePolicyOpts).Extract()
	if err != nil {
		return nil, err
	}
	t.Logf("Created Ike Policy: %s", ikePolicy.ID)

	return ikePolicy, nil
}

func deleteIkePolicy(t *testing.T, client *golangsdk.ServiceClient, ikePolicyId string) {
	t.Logf("Attempting to delete Ike Policy: %s", ikePolicyId)

	if err := ikepolicies.Delete(client, ikePolicyId).Err; err != nil {
		t.Fatalf("Unable to delete Ike Policy: %s", err)
	}

	t.Logf("Ike Policy is deleted: %s", ikePolicyId)
}

func updateIkePolicy(t *testing.T, client *golangsdk.ServiceClient, ikePolicy string) error {
	t.Logf("Attempting to update Ike policy")

	policyNewName := tools.RandomString("update-ike-", 8)

	updateOpts := ikepolicies.UpdateOpts{
		Name:          policyNewName,
		AuthAlgorithm: ikepolicies.AuthAlgorithm("sha1"),
	}

	if err := ikepolicies.Update(client, ikePolicy, updateOpts).Err; err != nil {
		return err
	}
	t.Logf("Ike Policy successfully updated: %s", ikePolicy)
	return nil
}
