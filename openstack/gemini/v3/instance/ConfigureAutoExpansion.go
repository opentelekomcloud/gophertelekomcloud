package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DiskAutoExpansionOpts struct {
	InstanceIds  []string                 `json:"instance_ids" required:"true"`
	SwitchOption string                   `json:"switch_option,omitempty"`
	Policy       *DiskAutoExpansionPolicy `json:"policy,omitempty"`
}

type DiskAutoExpansionPolicy struct {
	Threshold int `json:"threshold,omitempty"`
	Step      int `json:"step,omitempty"`
	Size      int `json:"size,omitempty"`
}

func ConfigureAutoExpansion(client *golangsdk.ServiceClient, opts DiskAutoExpansionOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", "disk-auto-expansion"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
