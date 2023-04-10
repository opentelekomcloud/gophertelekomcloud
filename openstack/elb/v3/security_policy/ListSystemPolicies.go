package security_policy

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListSystemPolicies(client *golangsdk.ServiceClient) ([]SystemPolicy, error) {
	raw, err := client.Get(client.ServiceURL("system-security-policies"), nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}

	var res []SystemPolicy

	err = extract.IntoSlicePtr(raw.Body, &res, "system_security_policies")
	return res, err
}

type SystemPolicy struct {
	ProjectId string `json:"project_id"`
	Name      string `json:"name"`
	Protocols string `json:"protocols"`
	Ciphers   string `json:"ciphers"`
}
