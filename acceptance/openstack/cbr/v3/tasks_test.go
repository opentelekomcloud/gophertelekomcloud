package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/tasks"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestTaskListAll(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	th.AssertNoErr(t, err)

	listOpts := tasks.ListOpts{}
	tasksList, err := tasks.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, task := range tasksList {
		tools.PrintResource(t, task)
	}
}
