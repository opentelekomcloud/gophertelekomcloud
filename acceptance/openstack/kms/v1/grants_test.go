package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/kms/v1/grants"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestKmsGrantsLifecycle(t *testing.T) {
	client, err := clients.NewKMSV1Client()
	th.AssertNoErr(t, err)

	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	if kmsID == "" {
		t.Skip("OS_KMS_ID env var is missing but KMSv1 grant test requires")
	}

	createOpts := grants.CreateOpts{
		KeyID:                kmsID,
		GranteePrincipal:     client.UserID,
		GranteePrincipalType: "user",
		Operations:           []string{"describe-key", "create-datakey", "encrypt-datakey"},
		Name:                 "my_grant",
	}
	createGrant, err := grants.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		deleteOpts := grants.DeleteOpts{
			KeyID:   kmsID,
			GrantID: createGrant.GrantID,
		}
		th.AssertNoErr(t, grants.Delete(client, deleteOpts).ExtractErr())
	}()

	listOpts := grants.ListOpts{
		KeyID: kmsID,
	}
	grantList, err := grants.List(client, listOpts).Extract()
	th.AssertNoErr(t, err)

	var found *grants.Grant
	for _, v := range grantList.Grants {
		if v.GrantID == createGrant.GrantID {
			found = &v
			break
		}
	}
	if found == nil {
		t.Fatal("created grant wasn't found by ID")
	}
	th.AssertEquals(t, createOpts.Name, found.Name)
	th.AssertEquals(t, createOpts.GranteePrincipal, found.GranteePrincipal)
	th.AssertEquals(t, len(createOpts.Operations), len(found.Operations))
}

func TestKmsGrantsDomainLifecycle(t *testing.T) {
	client, err := clients.NewKMSV1Client()
	th.AssertNoErr(t, err)

	kmsID := clients.EnvOS.GetEnv("KMS_ID")
	if kmsID == "" {
		t.Skip("OS_KMS_ID env var is missing but KMSv1 grant test requires")
	}

	createOpts := grants.CreateOpts{
		KeyID:                kmsID,
		GranteePrincipal:     client.DomainID,
		GranteePrincipalType: "domain",
		Operations:           []string{"describe-key", "create-datakey", "encrypt-datakey"},
		Name:                 "my_grant_domain",
	}
	createGrant, err := grants.Create(client, createOpts).Extract()
	th.AssertNoErr(t, err)

	defer func() {
		deleteOpts := grants.DeleteOpts{
			KeyID:   kmsID,
			GrantID: createGrant.GrantID,
		}
		th.AssertNoErr(t, grants.Delete(client, deleteOpts).ExtractErr())
	}()

	listOpts := grants.ListOpts{
		KeyID: kmsID,
	}
	grantList, err := grants.List(client, listOpts).Extract()
	th.AssertNoErr(t, err)

	var found *grants.Grant
	for _, v := range grantList.Grants {
		if v.GrantID == createGrant.GrantID {
			found = &v
			break
		}
	}
	if found == nil {
		t.Fatal("created grant wasn't found by ID")
	}
	th.AssertEquals(t, createOpts.Name, found.Name)
	th.AssertEquals(t, createOpts.GranteePrincipal, found.GranteePrincipal)
	th.AssertEquals(t, len(createOpts.Operations), len(found.Operations))
}
