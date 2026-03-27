package v2

import (
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	ns "github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/namespace"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/secret"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestSecretLifecycle(t *testing.T) {
	t.Skip("Tenant not whitelisted to run CCI")
	client, err := clients.NewCCIClient()
	th.AssertNoErr(t, err)

	nsName := "cci-namespace-" + strconv.Itoa(tools.RandomInt(1, 1000))
	secretName := "test-sn-" + strconv.Itoa(tools.RandomInt(1, 1000))

	createNsOpts := ns.CreateOpts{
		APIVersion: "cci/v2",
		Kind:       "Namespace",
		Metadata: ns.Metadata{
			Name: nsName,
		},
	}

	t.Logf("Attempting to create namespace")
	_, err = ns.Create(client, createNsOpts)
	th.AssertNoErr(t, err)

	err = waitForStatusActive(client, 600, nsName)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete namespace")
		nsDeleteOpts := ns.DeleteOpts{
			Name: nsName,
		}
		_, err = ns.Delete(client, nsDeleteOpts)
		th.AssertNoErr(t, err)
		err = waitForStatusDeleted(client, 500, nsName)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to create secret")

	secretOpts := secret.CreateOpts{
		Namespace:  nsName,
		APIVersion: "cci/v2",
		Kind:       "Secret",
		Data: map[string]string{
			"key": "eHh4Cg==",
		},
		Metadata: &secret.ObjectMeta{
			Name: secretName,
		},
	}

	secretResp, err := secret.Create(client, secretOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, secretResp.Metadata.Name, secretName)
	th.AssertEquals(t, secretResp.APIVersion, secretOpts.APIVersion)
	th.AssertEquals(t, secretResp.Data["key"], secretOpts.Data["key"])

	t.Cleanup(func() {
		t.Logf("Attempting to delete secret")
		secretDeleteOpts := secret.DeleteOpts{
			Namespace: nsName,
			Name:      secretName,
		}
		_, err = secret.Delete(client, secretDeleteOpts)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to update secret")

	updateSecretOpts := secret.UpdateOpts{
		APIVersion: "cci/v2",
		Data:       map[string]string{"key": "dGVzdA=="},
		Kind:       "Secret",
		Metadata: &secret.ObjectMeta{
			Name: secretName,
		},
	}

	updateResp, err := secret.Update(client, nsName, secretName, updateSecretOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, updateResp.Metadata.Name, secretName)
	th.AssertEquals(t, updateResp.Data["key"], updateSecretOpts.Data["key"])

	t.Logf("Attempting to get secret")

	getResp, err := secret.Get(client, nsName, secretName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getResp.Metadata.Name, secretName)
	th.AssertEquals(t, updateResp.Data["key"], getResp.Data["key"])

	t.Logf("Attempting to list secrets")
	secrets, err := secret.List(client, nsName, secret.ListOpts{})
	th.AssertNoErr(t, err)
	found := false
	for _, s := range secrets {
		if s.Metadata.Name == secretName {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)
}
