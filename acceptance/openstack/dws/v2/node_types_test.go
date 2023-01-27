package v1

import (
	"strings"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v2/cluster"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDWSNodeTypes(t *testing.T) {
	client, err := clients.NewDWSV1Client()
	th.AssertNoErr(t, err)
	client.ResourceBase = strings.Replace(client.ResourceBase, "v1.0/", "v2/", 1)

	nodes, err := cluster.ListNodeTypes(client)
	th.AssertNoErr(t, err)

	if len(nodes) > 0 {
		for _, node := range nodes {
			tools.PrintResource(t, node)
		}
	} else {
		t.Fatal("empty flavors list")
	}

}
