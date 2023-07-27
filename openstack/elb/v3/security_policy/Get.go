package security_policy

import (
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*SecurityPolicy, error) {
	raw, err := client.Get(client.ServiceURL("security-policies", id), nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	return extra(err, raw)
}

func extra(err error, raw *http.Response) (*SecurityPolicy, error) {
	if err != nil {
		return nil, err
	}

	var res SecurityPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "security_policy")
	return &res, err
}

type SecurityPolicy struct {
	// Specifies the ID of the custom security policy.
	Id string `json:"id"`
	// Specifies the project ID of the custom security policy.
	ProjectId string `json:"project_id"`
	// Specifies the name of the custom security policy.
	Name string `json:"name"`
	// Provides supplementary information about the custom security policy.
	Description string `json:"description"`
	// Specifies the listeners that use the custom security policies.
	Listeners []ListenerRef `json:"listeners"`
	// Lists the TLS protocols supported by the custom security policy.
	Protocols []string `json:"protocols"`
	// Lists the cipher suites supported by the custom security policy.
	Ciphers []string `json:"ciphers"`
	// Specifies the time when the custom security policy was created.
	CreatedAt string `json:"created_at"`
	// Specifies the time when the custom security policy was updated.
	UpdatedAt string `json:"updated_at"`
}

type ListenerRef struct {
	// Specifies the listener ID.
	Id string `json:"id"`
}
