package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Protocols   []string `json:"protocols" required:"true"`
	Ciphers     []string `json:"ciphers" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SecurityPolicy, error) {
	b, err := build.RequestBody(opts, "security_policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("security-policies"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	if err != nil {
		return nil, err
	}

	var res SecurityPolicy
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SecurityPolicy struct {
	SecurityPolicy PolicyRef `json:"security_policy"`
	RequestId      string    `json:"request_id"`
}

type PolicyRef struct {
	ID          string        `json:"id"`
	ProjectId   string        `json:"project_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Listeners   []ListenerRef `json:"listeners"`
	Protocols   []string      `json:"protocols"`
	Ciphers     []string      `json:"ciphers"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

type ListenerRef struct {
	ID string `json:"id"`
}
