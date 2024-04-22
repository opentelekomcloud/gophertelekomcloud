package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func DeleteResourceTag(client *golangsdk.ServiceClient, opts TagsActionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.DeleteWithBody(client.ServiceURL("functions", opts.Id, "tags", "delete"), b, nil)
	return err
}
