package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Address          string `json:"address" required:"true"`
	ProtocolPort     int    `json:"protocol_port" required:"true"`
	Name             string `json:"name,omitempty"`
	ProjectID        string `json:"project_id,omitempty"`
	Weight           *int   `json:"weight,omitempty"`
	SubnetID         string `json:"subnet_cidr_id,omitempty"`
	AdminStateUp     *bool  `json:"admin_state_up,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

func Create(client *golangsdk.ServiceClient, poolID string, opts CreateOpts) (*Member, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("pools", poolID, "members").Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "member")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Member Member `json:"member"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Member, nil
}
