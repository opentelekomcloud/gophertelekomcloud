package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResizeInstanceOpts struct {
	InstanceId         string          `json:"-"`
	SpecCode           string          `json:"spec_code" required:"true"`
	NewCapacity        float64         `json:"new_capacity" required:"true"`
	BssParam           DcsBssParamOpts `json:"bss_param,omitempty"`
	ReservedIp         []string        `json:"reserved_ip,omitempty"`
	ChangeType         string          `json:"change_type,omitempty"`
	AvailableZones     []string        `json:"available_zones,omitempty"`
	NodeList           []string        `json:"node_list,omitempty"`
	ExecuteImmediately *bool           `json:"execute_immediately"`
}

type DcsBssParamOpts struct {
	IsAutoPay string `json:"is_auto_pay,omitempty"`
}

func Resize(client *golangsdk.ServiceClient, opts ResizeInstanceOpts) (err error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("instances", opts.InstanceId, "resize"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
