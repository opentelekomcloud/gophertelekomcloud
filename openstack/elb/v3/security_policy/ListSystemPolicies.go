package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListSystemPolicies(client *golangsdk.ServiceClient) ([]SystemSecurityPolicy, error) {
	raw, err := client.Get(client.ServiceURL("system-security-policies"), nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}

	var res []SystemSecurityPolicy
	err = extract.IntoSlicePtr(raw.Body, &res, "system_security_policies")
	return res, err
}

type SystemSecurityPolicy struct {
	// Specifies the name of the system security policy.
	Name string `json:"name"`
	// Lists the TLS protocols supported by the system security policy.
	Protocols string `json:"protocols"`
	// Lists the cipher suites supported by the system security policy.
	Ciphers string `json:"ciphers"`
	// Specifies the project ID.
	ProjectId string `json:"project_id"`
}
