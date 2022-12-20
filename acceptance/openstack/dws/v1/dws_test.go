package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/cluster"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDWS(t *testing.T) {
	client, err := clients.NewDWSV1Client()
	th.AssertNoErr(t, err)

	newCluster, err := cluster.CreateCluster(client, cluster.CreateClusterOpts{
		NodeType:     "dws.m3.xlarge",
		NumberOfNode: 3,
	})
	th.AssertNoErr(t, err)
}
