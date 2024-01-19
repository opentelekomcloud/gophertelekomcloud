package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateLoginPolicyOpts struct {
	// Validity period (days) to disable users if they have not logged in within the period. Value range: 0-240.
	// Validity period (days) to disable users if they have not logged in within the period.
	// If this parameter is set to 0, no users will be disabled.
	AccountValidityPeriod *int `json:"account_validity_period,omitempty"`
	// Custom information that will be displayed upon successful login.
	CustomInfoForLogin string `json:"custom_info_for_login,omitempty"`
	// Duration (minutes) to lock users out.
	LockoutDuration int `json:"lockout_duration,omitempty"`
	// Number of unsuccessful login attempts to lock users out.
	LoginFailedTimes int `json:"login_failed_times,omitempty"`
	// Period (minutes) to count the number of unsuccessful login attempts.
	PeriodWithLoginFailures int `json:"period_with_login_failures,omitempty"`
	// Session timeout (minutes) that will apply if you or users created using your account
	// do not perform any operations within a specific period.
	SessionTimeout int `json:"session_timeout,omitempty"`
	// Indicates whether to display last login information upon successful login.
	ShowRecentLoginInfo *bool `json:"show_recent_login_info,omitempty"`
}

func UpdateLoginAuthPolicy(client *golangsdk.ServiceClient, id string, opts UpdateLoginPolicyOpts) (*LoginPolicy, error) {
	b, err := build.RequestBody(opts, "login_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/login-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "login-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LoginPolicy
	return &res, extract.IntoStructPtr(raw.Body, &res, "login_policy")
}
