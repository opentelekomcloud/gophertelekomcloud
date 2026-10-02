package rocketmq

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/rocketmq/instances"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRocketMQMonitoringDimensions(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	instanceID := prepareRocketMQInstance(t, client)

	dimensions, err := instances.GetMonitoringDimensions(client, instanceID)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, dimensions)

	th.AssertEquals(t, 1, len(dimensions.InstanceIDs))
	th.AssertEquals(t, instanceID, dimensions.InstanceIDs[0].Name)

	byName := make(map[string]instances.Dimension)
	for _, d := range dimensions.Dimensions {
		byName[d.Name] = d
	}
	for _, name := range []string{
		"reliablemq_instance_id",
		"reliablemq_broker",
		"reliablemq_topics",
		"reliablemq_groups",
		"reliablemq_dlq_topics",
	} {
		d, ok := byName[name]
		th.AssertEquals(t, true, ok)
		th.AssertEquals(t, true, len(d.Metrics) > 0)
	}
	th.AssertDeepEquals(t, []string{"instance_ids"}, byName["reliablemq_instance_id"].KeyName)
	th.AssertEquals(t, "reliablemq_groups_topics", byName["reliablemq_groups"].Children[0].Name)
}
