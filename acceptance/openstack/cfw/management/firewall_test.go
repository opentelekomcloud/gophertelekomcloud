package management

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	managementv1 "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v1/management"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCFWList(t *testing.T) {
	clientv1, err := clients.NewCFWV1Client()
	th.AssertNoErr(t, err)

	queryOpts := managementv1.ListOpts{
		Limit: 1024,
	}
	_, err = managementv1.List(clientv1, queryOpts)
	th.AssertNoErr(t, err)
}
