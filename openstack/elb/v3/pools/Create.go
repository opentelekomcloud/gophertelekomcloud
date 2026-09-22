package pools

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	LBMethod                 string              `json:"lb_algorithm" required:"true"`
	Protocol                 string              `json:"protocol" required:"true"`
	LoadbalancerID           string              `json:"loadbalancer_id,omitempty"`
	ListenerID               string              `json:"listener_id,omitempty"`
	ProjectID                string              `json:"project_id,omitempty"`
	Name                     string              `json:"name,omitempty"`
	Description              string              `json:"description,omitempty"`
	Persistence              *SessionPersistence `json:"session_persistence,omitempty"`
	SlowStart                *SlowStart          `json:"slow_start,omitempty"`
	AdminStateUp             *bool               `json:"admin_state_up,omitempty"`
	DeletionProtectionEnable *bool               `json:"member_deletion_protection_enable,omitempty"`
	VpcId                    string              `json:"vpc_id,omitempty"`
	Type                     string              `json:"type,omitempty"`
	ProtectionStatus         string              `json:"protection_status,omitempty"`
	ProtectionReason         string              `json:"protection_reason,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Pool, error) {
	b, err := build.RequestBody(opts, "pool")
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("pools"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
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
