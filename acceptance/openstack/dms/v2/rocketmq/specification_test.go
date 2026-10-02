package rocketmq

import (
	"fmt"
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/instances"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/specification"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// TestRocketMQSpecification always creates its own instance:
// storage expansion cannot be reverted.
func TestRocketMQSpecification(t *testing.T) {
	if clients.EnvOS.GetEnv("RUN_DMS_ROCKETMQ") == "" {
		t.Skip("OS_RUN_DMS_ROCKETMQ env var is missing, RocketMQ instance creation takes ~30 minutes")
	}

	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := createRocketMQInstance(t, client)
	t.Cleanup(func() { deleteRocketMQInstance(t, client, instanceID) })

	extendProducts, err := specification.ListProducts(client, instanceID, specification.ListProductsOpts{Limit: 10})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, extendProducts)
	th.AssertEquals(t, "rocketmq", extendProducts.Engine)
	th.AssertEquals(t, true, len(extendProducts.Products) > 0)
	th.AssertEquals(t, len(extendProducts.Products), extendProducts.Total)

	instance, err := instances.Get(client, instanceID)
	th.AssertNoErr(t, err)
	newStorage := instance.StorageSpace + 100

	t.Logf("Attempting to expand storage of RocketMQ instance %s to %d GB", instanceID, newStorage)
	resized, err := specification.Resize(client, instanceID, specification.ResizeOpts{
		OperType:        "storage",
		NewStorageSpace: newStorage,
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, true, resized.JobID != "")

	var lastState string
	th.AssertNoErr(t, golangsdk.WaitFor(1800, func() (bool, error) {
		instance, err = instances.Get(client, instanceID)
		if err != nil {
			return false, err
		}
		if state := fmt.Sprintf("status=%s storage=%d", instance.Status, instance.StorageSpace); state != lastState {
			t.Logf("RocketMQ instance %s: %s", instanceID, state)
			lastState = state
		}
		return instance.StorageSpace == newStorage && instance.Status == rocketMQTargetStatus, nil
	}))
	th.AssertEquals(t, newStorage, instance.TotalStorageSpace)
}
