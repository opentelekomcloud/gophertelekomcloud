package rocketmq

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/products"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRocketMQProductsList(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	page, err := products.List(client, products.ListOpts{Engine: rocketMQEngine, Limit: 1})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(page.Products))
	th.AssertEquals(t, true, page.Total > 1)
	th.AssertEquals(t, 1, page.NextOffset)
	th.AssertEquals(t, -1, page.PreviousOffset)

	next, err := products.List(client, products.ListOpts{Engine: rocketMQEngine, Limit: 1, Offset: 1})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 1, len(next.Products))
	th.AssertEquals(t, 0, next.PreviousOffset)
	th.AssertEquals(t, true, next.Products[0].ProductId != page.Products[0].ProductId)

	product := getRocketMQProduct(t, client)
	single, err := products.List(client, products.ListOpts{Engine: rocketMQEngine, ProductId: product.ProductId})
	th.AssertNoErr(t, err)
	tools.PrintResource(t, single)
	th.AssertEquals(t, 1, single.Total)
	th.AssertEquals(t, 1, len(single.Products))

	properties := single.Products[0].Properties
	th.AssertEquals(t, rocketMQEngineVersion, properties.EngineVersions)
	th.AssertEquals(t, product.ProductId, properties.ProductAlias)
	th.AssertEquals(t, "1", properties.BrokerNum)
	th.AssertEquals(t, true, properties.MaxTopic != "")
	th.AssertEquals(t, true, properties.MaxConsumer != "")
	th.AssertEquals(t, true, properties.RCU != "")
	th.AssertEquals(t, true, properties.MaxTPSPerRCU != "")
	th.AssertEquals(t, true, properties.MaxStorage != "")
	th.AssertEquals(t, true, properties.MinStorage != "")
}
