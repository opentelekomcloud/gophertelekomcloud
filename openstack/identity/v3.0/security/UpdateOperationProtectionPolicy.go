package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateProtectionPolicyOpts struct {
	// Indicates whether operation protection has been enabled. The value can be true or false.
	OperationProtection *bool `json:"operation_protection" required:"true"`
	// Specifies the IAM attributes which user can modify
	AllowUser *AllowUser `json:"allow_user,omitempty"`
	// Specifies whether a person is designated for verification.
	// Valid options are the on and off.
	AdminCheck string `json:"admin_check,omitempty"`
	// Specifies mobile number used for verification
	Mobile string `json:"mobile,omitempty"`
	// Specifies email address used for verification
	Email string `json:"email,omitempty"`
	// Specifies the verification method. This parameter is mandatory when admin_check is set to on.
	// The valid options are mobile and email.
	Scene string `json:"scene,omitempty"`
}

type AllowUser struct {
	// Specifies whether IAM users are allowed to manage access keys.
	ManageAccessKey *bool `json:"manage_accesskey,omitempty"`
	// Specifies whether IAM users are allowed to change their email addresses.
	ManageEmail *bool `json:"manage_email,omitempty"`
	// Specifies whether IAM users are allowed to change their mobile numbers.
	ManageMobile *bool `json:"manage_mobile,omitempty"`
	// Specifies whether IAM users are allowed to change their passwords.
	ManagePassword *bool `json:"manage_password,omitempty"`
}

func UpdateOperationProtectionPolicy(client *golangsdk.ServiceClient, id string, opts UpdateProtectionPolicyOpts) (*ProtectionPolicy, error) {
	b, err := build.RequestBody(opts, "protect_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/protect-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "protect-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ProtectionPolicy
	return &res, extract.IntoStructPtr(raw.Body, &res, "protect_policy")
}
