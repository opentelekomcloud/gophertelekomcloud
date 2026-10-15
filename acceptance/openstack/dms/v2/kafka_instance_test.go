package v2

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/instances/lifecycle"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

// kafkaInstanceName returns the name of the long-living Kafka instance from OS_DMS_KAFKA_INSTANCE_NAME
// and skips the test if it is unset.
func kafkaInstanceName(t *testing.T) string {
	t.Helper()
	name := clients.EnvOS.GetEnv("DMS_KAFKA_INSTANCE_NAME")
	if name == "" {
		t.Skip("OS_DMS_KAFKA_INSTANCE_NAME is not set")
	}
	return name
}

// existingKafkaInstance returns a long-living Kafka instance that tests must not delete or reconfigure permanently.
// The test is skipped if the instance is not found.
func existingKafkaInstance(t *testing.T, client *golangsdk.ServiceClient, name string) *lifecycle.Instance {
	t.Helper()

	list, err := lifecycle.List(client, lifecycle.ListOpts{
		Engine:         dmsEngine,
		Name:           name,
		ExactMatchName: "true",
	})
	th.AssertNoErr(t, err)
	if len(list.Instances) == 0 {
		t.Skipf("Kafka instance %q not found in this project", name)
	}
	th.AssertEquals(t, 1, len(list.Instances))
	return &list.Instances[0]
}

func TestExistingKafkaInstanceInfo(t *testing.T) {
	name := kafkaInstanceName(t)

	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instance := existingKafkaInstance(t, client, name)
	tools.PrintResource(t, instance)

	th.AssertEquals(t, dmsEngine, instance.Engine)
	th.AssertEquals(t, dmsTargetStatus, instance.Status)
}
