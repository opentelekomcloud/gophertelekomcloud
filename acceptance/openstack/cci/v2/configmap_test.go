package v2

import (
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/configmap"
	ns "github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/namespace"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestConfigMapLifecycle(t *testing.T) {
	t.Skip("Tenant not whitelisted to run CCI")
	client, err := clients.NewCCIClient()
	th.AssertNoErr(t, err)

	nsName := "cci-namespace-" + strconv.Itoa(tools.RandomInt(1, 1000))
	configMapName := "test-cm-" + strconv.Itoa(tools.RandomInt(1, 1000))

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

	t.Logf("Attempting to create configmap")
	createOpts := configmap.CreateOpts{
		APIVersion: "cci/v2",
		Kind:       "ConfigMap",
		Data: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		Metadata: configmap.ObjectMeta{
			Name: configMapName,
			Labels: map[string]string{
				"usage": "just-for-test",
			},
		},
	}

	cm, err := configmap.Create(client, nsName, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, configMapName, cm.Metadata.Name)
	th.AssertEquals(t, "value1", cm.Data["key1"])

	t.Cleanup(func() {
		t.Logf("Attempting to delete configmap")
		deleteOpts := configmap.DeleteOpts{
			PropagationPolicy: "Background",
		}
		deletedCm, err := configmap.Delete(client, nsName, configMapName, deleteOpts)
		th.AssertNoErr(t, err)
		th.AssertEquals(t, configMapName, deletedCm.Metadata.Name)
	})

	t.Logf("Attempting to list configmaps")
	listOpts := configmap.ListOpts{}
	cms, err := configmap.List(client, nsName, listOpts)
	th.AssertNoErr(t, err)
	found := false
	for _, item := range cms {
		if item.Metadata.Name == configMapName {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)

	t.Logf("Attempting to get configmap")
	getCm, err := configmap.Get(client, nsName, configMapName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, configMapName, getCm.Metadata.Name)
	th.AssertEquals(t, "value1", getCm.Data["key1"])

	t.Logf("Attempting to update configmap")
	updateOpts := configmap.UpdateOpts{
		APIVersion: "cci/v2",
		Kind:       "ConfigMap",
		Data: map[string]string{
			"key1": "updated-value1",
			"key2": "updated-value2",
		},
		Metadata: &configmap.ObjectMeta{
			Name: configMapName,
			Labels: map[string]string{
				"usage": "updated-test",
			},
		},
	}

	updatedCm, err := configmap.Update(client, nsName, configMapName, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "updated-value1", updatedCm.Data["key1"])
	th.AssertEquals(t, "updated-test", updatedCm.Metadata.Labels["usage"])
}
