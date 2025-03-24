package v2

import (
	"fmt"
	"strconv"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	ns "github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/namespace"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestNamespaceLifecycle(t *testing.T) {
	t.Skip("Tenant not whitelisted to run CCI")
	client, err := clients.NewCCIClient()
	th.AssertNoErr(t, err)

	nsName := "cci-namespace-" + strconv.Itoa(tools.RandomInt(1, 1000))

	createOpts := ns.CreateOpts{
		APIVersion: "cci/v2",
		Kind:       "Namespace",
		Metadata: ns.Metadata{
			Name: nsName,
		},
	}

	t.Logf("Attempting to create namespace")
	namespace, err := ns.Create(client, createOpts)
	th.AssertNoErr(t, err)

	err = waitForStatusActive(client, 600, namespace.Metadata.Name)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete namespace")
		deleteOpts := ns.DeleteOpts{
			Name: createOpts.Metadata.Name,
		}
		_, err = ns.Delete(client, deleteOpts)
		th.AssertNoErr(t, err)
		err = waitForStatusDeleted(client, 500, nsName)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to retrieve namespace")
	nGet, err := ns.Get(client, createOpts.Metadata.Name)
	th.AssertNoErr(t, err)

	th.AssertEquals(t, createOpts.Metadata.Name, nGet.Metadata.Name)
	th.AssertEquals(t, createOpts.Kind, nGet.Kind)

	listOpts := ns.ListOpts{}
	t.Logf("Attempting to retrieve all namespaces")
	nList, err := ns.List(client, listOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Metadata.Name, nList[0].Metadata.Name)
}

func waitForStatusActive(client *golangsdk.ServiceClient, secs int, name string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		namespace, err := ns.Get(client, name)
		if err != nil {
			return false, err
		}

		if namespace.Status.Phase == "Active" {
			return true, nil
		}
		if namespace.Status.Phase == "Failed" {
			err = fmt.Errorf("Creation failed %s.\n", namespace.Status.Phase)
			return false, err
		}

		return false, nil
	})
}

func waitForStatusDeleted(client *golangsdk.ServiceClient, secs int, name string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		namespace, err := ns.Get(client, name)
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return true, nil
			}
			return false, err
		}

		if namespace.Status.Phase == "Failed" {
			err = fmt.Errorf("Creation failed %s.\n", namespace.Status.Phase)
			return false, err
		}

		return false, nil
	})
}
