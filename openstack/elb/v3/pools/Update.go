package pools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Name                     *string             `json:"name,omitempty"`
	Description              *string             `json:"description,omitempty"`
	LBMethod                 string              `json:"lb_algorithm,omitempty"`
	Persistence              *SessionPersistence `json:"session_persistence,omitempty"`
	AdminStateUp             *bool               `json:"admin_state_up,omitempty"`
	SlowStart                *SlowStart          `json:"slow_start,omitempty"`
	DeletionProtectionEnable *bool               `json:"member_deletion_protection_enable,omitempty"`
	VpcId                    string              `json:"vpc_id,omitempty"`
	Type                     string              `json:"type,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Pool, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("pools", id).Build()
	if err != nil {
		return nil, err
	}
	b, err := build.RequestBody(opts, "pool")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Pool Pool `json:"pool"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Pool, nil
}
