package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetLoginAuthPolicy(client *golangsdk.ServiceClient, id string) (*LoginPolicy, error) {
	// GET  /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/login-policy
	raw, err := client.Get(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "login-policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res LoginPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "login_policy")
	return &res, err
}

type LoginPolicy struct {
	// Validity period (days) to disable users if they have not logged in within the period. Value range: 0-240.
	// Validity period (days) to disable users if they have not logged in within the period.
	// If this parameter is set to 0, no users will be disabled.
	AccountValidityPeriod *int `json:"account_validity_period"`
	// Custom information that will be displayed upon successful login.
	CustomInfoForLogin string `json:"custom_info_for_login"`
	// Duration (minutes) to lock users out.
	LockoutDuration int `json:"lockout_duration"`
	// Number of unsuccessful login attempts to lock users out.
	LoginFailedTimes int `json:"login_failed_times"`
	// Period (minutes) to count the number of unsuccessful login attempts.
	PeriodWithLoginFailures int `json:"period_with_login_failures"`
	// Session timeout (minutes) that will apply if you or users created using your account
	// do not perform any operations within a specific period.
	SessionTimeout int `json:"session_timeout"`
	// Indicates whether to display last login information upon successful login.
	ShowRecentLoginInfo *bool `json:"show_recent_login_info"`
}
