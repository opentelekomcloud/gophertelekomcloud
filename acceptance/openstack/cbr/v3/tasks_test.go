package v3

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cbr/v3/tasks"
)

func TestTaskListAll(t *testing.T) {
	client, err := clients.NewCbrV3Client()
	if err != nil {
		t.Fatalf("Unable to create a CBRv3 client: %s", err)
	}

	listOpts := tasks.ListOpts{}
	pages, err := tasks.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to fetch CBR tasks pages: %s", err)
	}
	tasksList, err := tasks.ExtractTasks(pages)
	if err != nil {
		t.Fatalf("Unable to extract CBR tasks pages: %s", err)
	}

	for _, task := range tasksList {
		tools.PrintResource(t, task)
	}
}
