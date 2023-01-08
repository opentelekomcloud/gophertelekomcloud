package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type AddNodeOpts struct {
	Type       string      `json:"type" required:"true"`
	SpecCode   string      `json:"spec_code" required:"true"`
	Num        int         `json:"num" required:"true"`
	Volume     *VolumeNode `json:"volume,omitempty"`
	InstanceId string      `json:"-"`
}

type VolumeNode struct {
	Size int `json:"size"`
}

func AddNode(client *golangsdk.ServiceClient, opts AddNodeOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "enlarge"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return extractJob(err, raw)
}
