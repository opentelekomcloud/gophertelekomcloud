package v3

import (
	"os"
	"reflect"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/agency"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/policies"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPolicyLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV3AdminClient() to be initialized.")
	}

	client, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	newPolicy := createPolicy(t, client)
	th.AssertNoErr(t, err)

	defer func() {
		err = policies.Delete(client, newPolicy.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	listOpts := policies.ListOpts{
		DisplayName: newPolicy.DisplayName,
	}

	allPages, err := policies.List(client, listOpts)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, newPolicy.Description, allPages.Roles[0].Description)
	th.AssertEquals(t, newPolicy.Type, allPages.Roles[0].Type)
	th.AssertEquals(t, newPolicy.DisplayName, allPages.Roles[0].DisplayName)
	th.AssertEquals(t, newPolicy.Policy.Version, allPages.Roles[0].Policy.Version)
	if !reflect.DeepEqual(newPolicy.Policy.Statement, allPages.Roles[0].Policy.Statement) {
		t.Error("Statement parameters are different")
	}

	newName := tools.RandomString("policy-test-new-", 5)
	newDescription := tools.RandomString("description-new-", 5)

	updateOpts := policies.CreateOpts{
		DisplayName: newName,
		Type:        "AX",
		Description: newDescription,
		Policy: policies.CreatePolicy{
			Version: "1.1",
			Statement: []policies.CreateStatement{
				{Action: []string{
					"obs:bucket:ListBucket",
				},
					Effect: "Deny",
					Condition: map[string]map[string][]string{
						"StringNotStartWithIfExists": {
							"g:ServiceName": []string{"ht"},
						},
						"StringNotEndWith": {
							"g:UserName": []string{"3"},
						},
					},
					Resource: []string{"OBS:*:*:bucket:*"},
				},
			},
		},
	}

	updatePolicy, err := policies.Update(client, newPolicy.ID, updateOpts).Extract()

	th.AssertNoErr(t, err)

	getPolicy, err := policies.Get(client, updatePolicy.ID).Extract()

	th.AssertNoErr(t, err)

	th.AssertEquals(t, updatePolicy.DisplayName, newName)
	th.AssertEquals(t, updatePolicy.Description, newDescription)
	th.AssertEquals(t, updatePolicy.Name, getPolicy.Name)
	th.AssertEquals(t, updatePolicy.Description, getPolicy.Description)
	th.AssertEquals(t, updatePolicy.Type, getPolicy.Type)
	th.AssertEquals(t, updatePolicy.DisplayName, getPolicy.DisplayName)
	th.AssertEquals(t, updatePolicy.Policy.Version, getPolicy.Policy.Version)
	if !reflect.DeepEqual(updatePolicy.Policy.Statement, getPolicy.Policy.Statement) {
		t.Error("Statement parameters are different")
	}
}

func TestAgencyPolicyLifecycle(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV30AdminClient() to be initialized.")
	}

	if os.Getenv("DELEGATED_DOMAIN") == "" {
		t.Skip("Unable to continue test without provided delegated domain name.")
	}

	client, err := clients.NewIdentityV30AdminClient()

	th.AssertNoErr(t, err)

	agencyOpts := agency.CreateOpts{
		Name:            tools.RandomString("test-agency-", 5),
		DomainID:        client.DomainID,
		DelegatedDomain: os.Getenv("DELEGATED_DOMAIN"),
	}

	createAgency, err := agency.Create(client, agencyOpts).Extract()

	th.AssertNoErr(t, err)

	defer func() {
		err = agency.Delete(client, createAgency.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	// this parameter is hardcoded for every custom policy made through agency
	actionParameter := "iam:agencies:assume"
	createOpts := policies.CreateOpts{
		DisplayName: tools.RandomString("policy-test-", 5),
		Type:        "AX",
		Description: tools.RandomString("Description", 5),
		Policy: policies.CreatePolicy{
			Version: "1.1",
			Statement: []policies.CreateStatement{
				{Action: []string{
					actionParameter,
				},
					Effect: "Allow",
					Resource: map[string][]string{
						"uri": {"/iam/agencies/" + createAgency.ID},
					},
				},
			},
		},
	}

	newPolicy, err := policies.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		err = policies.Delete(client, newPolicy.ID).ExtractErr()
		th.AssertNoErr(t, err)
	}()

	listOpts := policies.ListOpts{
		ID:   newPolicy.ID,
		Type: newPolicy.Type,
	}

	allPages, err := policies.List(client, listOpts)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, newPolicy.ID, allPages.Roles[0].ID)
	th.AssertEquals(t, newPolicy.Description, allPages.Roles[0].Description)
	th.AssertEquals(t, newPolicy.Type, allPages.Roles[0].Type)
	th.AssertEquals(t, newPolicy.DisplayName, allPages.Roles[0].DisplayName)
	th.AssertEquals(t, newPolicy.Policy.Version, allPages.Roles[0].Policy.Version)
	if !reflect.DeepEqual(newPolicy.Policy.Statement, allPages.Roles[0].Policy.Statement) {
		t.Error("Statement parameters are different")
	}

	newName := tools.RandomString("policy-test-new-", 5)
	newDescription := tools.RandomString("description-new-", 5)

	updateOpts := policies.CreateOpts{
		DisplayName: newName,
		Type:        "AX",
		Description: newDescription,
		Policy: policies.CreatePolicy{
			Version: "1.1",
			Statement: []policies.CreateStatement{
				{Action: []string{
					actionParameter,
				},
					Effect: "Allow",
					Resource: map[string][]string{
						"uri": {"/iam/agencies/" + createAgency.ID},
					},
				},
			},
		},
	}

	updatePolicy, err := policies.Update(client, newPolicy.ID, updateOpts).Extract()

	th.AssertNoErr(t, err)

	getPolicy, err := policies.Get(client, updatePolicy.ID).Extract()

	th.AssertNoErr(t, err)

	th.AssertEquals(t, updatePolicy.DisplayName, newName)
	th.AssertEquals(t, updatePolicy.Description, newDescription)
	th.AssertEquals(t, updatePolicy.Name, getPolicy.Name)
	th.AssertEquals(t, updatePolicy.Description, getPolicy.Description)
	th.AssertEquals(t, updatePolicy.Type, getPolicy.Type)
	th.AssertEquals(t, updatePolicy.DisplayName, getPolicy.DisplayName)
	th.AssertEquals(t, updatePolicy.Policy.Version, getPolicy.Policy.Version)
	if !reflect.DeepEqual(updatePolicy.Policy.Statement, getPolicy.Policy.Statement) {
		t.Error("Statement parameters are different")
	}
}

func TestList(t *testing.T) {
	if os.Getenv("OS_TENANT_ADMIN") == "" {
		t.Skip("Policy doesn't allow NewIdentityV30AdminClient() to be initialized.")
	}

	client, err := clients.NewIdentityV30AdminClient()
	th.AssertNoErr(t, err)

	listOpts := policies.ListOpts{}

	allPages, err := policies.List(client, listOpts)
	th.AssertNoErr(t, err)
	if allPages.Links.Self == "" {
		t.Error("Link parameter unmarshalled improperly")
	}
}

func createPolicy(t *testing.T, client *golangsdk.ServiceClient) *policies.Policy {
	createOpts := policies.CreateOpts{
		DisplayName: tools.RandomString("policy-test-", 5),
		Type:        "AX",
		Description: tools.RandomString("Description-", 5),
		Policy: policies.CreatePolicy{
			Version: "1.1",
			Statement: []policies.CreateStatement{
				{Action: []string{
					"obs:bucket:ListBucket",
				},
					Effect: "Allow",
					Condition: map[string]map[string][]string{
						"StringLikeAnyOfIfExists": {
							"obs:prefix": []string{"ht"},
						},
						"StringNotEndWith": {
							"g:UserName": []string{tools.RandomString("end_of_name-", 5)},
						},
					},
					Resource: []string{"OBS:*:*:bucket:some-bucket", "OBS:*:*:bucket:test-bucket"},
				},
				{Action: []string{
					"obs:bucket:ListBucket",
				},
					Effect: "Deny",
					Condition: map[string]map[string][]string{
						"StringEquals": {
							"g:UserId": []string{tools.RandomString("dummy_id-", 5)},
						},
					},
					Resource: []string{"OBS:*:*:bucket:bucket"},
				},
			},
		},
	}

	newPolicy, err := policies.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)
	return newPolicy
}
