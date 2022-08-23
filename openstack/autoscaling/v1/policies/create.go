package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// CreateOpts is a struct which will be used to create a policy
type CreateOpts struct {
	Name           string             `json:"scaling_policy_name" required:"true"`
	ID             string             `json:"scaling_group_id" required:"true"`
	Type           string             `json:"scaling_policy_type" required:"true"`
	AlarmID        string             `json:"alarm_id,omitempty"`
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	Action         ActionOpts         `json:"scaling_policy_action,omitempty"`
	CoolDownTime   int                `json:"cool_down_time,omitempty"`
}

// Create is a method which can be able to access to create the policy of autoscaling
// service.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(client.ServiceURL("scaling_policy"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
