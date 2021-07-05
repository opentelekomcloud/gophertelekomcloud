package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder is an interface by which can serialize the create parameters
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// CreateOpts is a struct which will be used to create a policy
type CreateOpts struct {
	PolicyName     string             `json:"scaling_policy_name" required:"true"`
	PolicyType     string             `json:"scaling_policy_type" required:"true"`
	ResourceID     string             `json:"scaling_resource_id" required:"true"`
	ResourceType   string             `json:"scaling_resource_type" required:"true"`
	AlarmID        string             `json:"alarm_id,omitempty"`
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	PolicyAction   ActionOpts         `json:"scaling_policy_action,omitempty"`
	CoolDownTime   int                `json:"cool_down_time,omitempty"`
}

type SchedulePolicyOpts struct {
	LaunchTime      string `json:"launch_time" required:"true"`
	RecurrenceType  string `json:"recurrence_type,omitempty"`
	RecurrenceValue string `json:"recurrence_value,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
}

type ActionOpts struct {
	Operation  string `json:"operation,omitempty"`
	Size       int    `json:"size,omitempty"`
	Percentage int    `json:"percentage,omitempty"`
	Limits     int    `json:"limits,omitempty"`
}

func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create is a method which can be able to access to create the policy of autoscaling
// service.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// UpdateOptsBuilder is an interface which can build the map parameter of update function
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	PolicyName     string             `json:"scaling_policy_name,omitempty"`
	PolicyType     string             `json:"scaling_policy_type,omitempty"`
	ResourceID     string             `json:"scaling_resource_id,omitempty"`
	ResourceType   string             `json:"scaling_resource_type,omitempty"`
	AlarmID        string             `json:"alarm_id,omitempty"`
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	Action         ActionOpts         `json:"scaling_policy_action,omitempty"`
	CoolDownTime   int                `json:"cool_down_time,omitempty"`
}

func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update is a method which can be able to update the policy via accessing to the
// autoscaling service with Put method and parameters
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	body, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Put(singleURL(client, id), body, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete is a method which can be able to access to delete a policy of autoscaling
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(singleURL(client, id), nil)
	return
}

// Get is a method which can be able to access to get a policy detailed information
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(singleURL(client, id), &r.Body, nil)
	return
}
