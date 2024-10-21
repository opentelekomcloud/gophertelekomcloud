package v2

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2/products"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestProductsList(t *testing.T) {
	client, err := clients.NewDmsV2Client()
	th.AssertNoErr(t, err)

	pd, err := products.List(client, products.ListOpts{Engine: dmsEngine, ProductId: "c6.2u4g.single"})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, len(pd.Products), 1)
}
