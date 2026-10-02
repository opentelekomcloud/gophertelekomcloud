package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/quotas"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestQuotasGet(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	all, err := quotas.Get(client, quotas.GetOpts{})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, all)
	types := make(map[string]quotas.Resource)
	for _, r := range all.Resources {
		types[r.Type] = r
	}
	for _, typ := range []string{"kafkaInstance", "rocketmqInstance", "tags"} {
		r, ok := types[typ]
		th.AssertEquals(t, true, ok)
		th.AssertEquals(t, true, r.Quota > 0)
	}

	rocketMQ, err := quotas.Get(client, quotas.GetOpts{
		IncludeTagsQuota: pointerto.Bool(false),
		OnlyQuota:        "reliability",
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(rocketMQ.Resources))
	th.AssertEquals(t, "rocketmqInstance", rocketMQ.Resources[0].Type)
	th.AssertEquals(t, types["rocketmqInstance"].Quota, rocketMQ.Resources[0].Quota)
}
