package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dis/v2/apps"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dis/v2/streams"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDIS(t *testing.T) {
	client, err := clients.NewDisV2Client()
	th.AssertNoErr(t, err)

	appName := tools.RandomString("app-create-test-", 3)
	err = apps.CreateApp(client, apps.CreateAppOpts{
		AppName: appName,
	})
	th.AssertNoErr(t, err)

	app, err := apps.DescribeApp(client, appName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, app)

	streamName := tools.RandomString("stream-create-test-", 3)
	err = streams.CreateStream(client, streams.CreateStreamOpts{
		StreamName:     streamName,
		PartitionCount: 3,
	})
	th.AssertNoErr(t, err)

	stream, err := streams.DescribeStream(client, streams.DescribeStreamOpts{
		StreamName: streamName,
	})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, stream)

	err = streams.CreatePolicyRule(client, streams.CreatePolicyRuleOpts{
		StreamName:    streamName,
		StreamId:      stream.StreamId,
		PrincipalName: client.DomainID,
		ActionType:    "putRecords",
		Effect:        "effect",
	})
	th.AssertNoErr(t, err)

	rule, err := streams.GetPolicyRule(client, streamName)
	th.AssertNoErr(t, err)
	tools.PrintResource(t, rule)
}
