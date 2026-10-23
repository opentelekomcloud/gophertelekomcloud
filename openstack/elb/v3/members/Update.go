package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name             *string `json:"name,omitempty"`
	Weight           *int    `json:"weight,omitempty"`
	AdminStateUp     *bool   `json:"admin_state_up,omitempty"`
	AvailabilityZone string  `json:"availability_zone,omitempty"`
}

func Update(client *golangsdk.ServiceClient, poolID, memberID string, opts UpdateOpts) (*Member, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("pools", poolID, "members", memberID).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "member")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
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
