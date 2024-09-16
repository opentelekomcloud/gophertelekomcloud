package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Name specifies the policy name. The value consists of 1 to 64 characters
	// and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name"`
	// OperationDefinition - scheduling configuration
	OperationDefinition *PolicyODCreate `json:"operation_definition"`
	// Enabled - whether to enable the policy, default: true
	Enabled *bool `json:"enabled,omitempty"`
	// OperationType - policy type
	OperationType OperationType `json:"operation_type"`
	// Trigger - time rule for the policy execution
	Trigger *Trigger `json:"trigger"`
}

type Trigger struct {
	Properties TriggerProperties `json:"properties"`
}

type TriggerProperties struct {
	// Pattern - Scheduling policy of the scheduler. Can't be empty.
	Pattern []string `json:"pattern"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Policy, error) {
	b, err := build.RequestBody(opts, "policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Policy
	err = extract.IntoStructPtr(raw.Body, &res, "policy")
	return &res, err
}
