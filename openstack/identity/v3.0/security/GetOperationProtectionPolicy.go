package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ProtectionPolicy struct {
	// Indicates whether operation protection has been enabled. The value can be true or false.
	OperationProtection *bool `json:"operation_protection"`
	// Specifies whether a person is designated for verification.
	AdminCheck string `json:"admin_check"`
	// The verification method
	Scene string `json:"scene"`
	// The IAM attributes which user can modify
	AllowUser *AllowUser `json:"allow_user"`
	// Specifies mobile number used for verification
	Mobile string `json:"mobile"`
	// Specifies email address used for verification
	Email string `json:"email"`
}

func GetOperationProtectionPolicy(client *golangsdk.ServiceClient, id string) (*ProtectionPolicy, error) {
	// GET /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/protect-policy
	raw, err := client.Get(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "protect-policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ProtectionPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "protect_policy")
	return &res, err
}
