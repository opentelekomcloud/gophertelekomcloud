package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/cluster"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dws/v1/snapshot"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestDWS(t *testing.T) {
	client, err := clients.NewDWSV1Client()
	th.AssertNoErr(t, err)

	aCluster, err := cluster.CreateCluster(client, cluster.CreateClusterOpts{})
	th.AssertNoErr(t, err)

	aSnapshot, err := snapshot.CreateSnapshot(client, snapshot.Snapshot{})
	th.AssertNoErr(t, err)

}
