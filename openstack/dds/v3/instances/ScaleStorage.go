package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ScaleStorageOpt struct {
	GroupId    string `json:"group_id,omitempty"`
	Size       string `json:"size" required:"true"`
	InstanceId string `json:"-"`
}

func ScaleStorage(client *golangsdk.ServiceClient, opts ScaleStorageOpt) (*string, error) {
	b, err := build.RequestBody(opts, "volume")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "enlarge-volume"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extractJob(err, raw)
}
