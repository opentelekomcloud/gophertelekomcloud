package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateACLRuleOpts struct {
	// Rule name.
	Name string `json:"name,omitempty"`
	// Internet protocol type of an address (0: IPv4, 1: IPv6).
	AddressType int `json:"address_type,omitempty"`
	// Rule action (0: Permit, 1: Deny).
	ActionType int `json:"action_type,omitempty"`
	// Rule status (0: Disabled, 1: Enabled)
	Status int `json:"status,omitempty"`
	// Rule application list. Rule application type:
	// HTTP, HTTPS, TLS1, DNS, SSH, MYSQL, SMTP, RDP, RDPS, VNC, POP3, IMAP4, SMTPS, POP3S, FTPS, ANY, or BGP.
	Applications []string `json:"applications,omitempty"`
	// JSON string converted from the applications field in the application list.
	ApplicationsJsonString string `json:"applicationsJsonString,omitempty"`
	// Persistent connection duration.
	LongConnectTime int64 `json:"long_connect_time,omitempty"`
	// Persistent connection duration (hour).
	LongConnectTimeHour int64 `json:"long_connect_time_hour,omitempty"`
	// Persistent connection duration (minutes).
	LongConnectTimeMinute int64 `json:"long_connect_time_minute,omitempty"`
	// Persistent connection duration (seconds).
	LongConnectTimeSecond int64 `json:"long_connect_time_second,omitempty"`
	// Whether to enable long connection (0: No, 1: Yes).
	LongConnectEnable int `json:"long_connect_enable,omitempty"`
	// Description of the rule.
	Description string `json:"description,omitempty"`
	// Direction of rule (0: Inbound, 1: Outbound). This parameter is mandatory when type is set to 0 (Internet rule) or 2 (NAT rule).
	Direction int `json:"direction,omitempty"`
	// Source address Data Transport Object.
	Source RuleAddressDtoRequest `json:"source,omitempty"`
	// Destination address Data Transport Object.
	Destination RuleAddressDtoRequest `json:"destination,omitempty"`
	// Service object associated with the rule.
	Service RuleServiceDto `json:"service,omitempty"`
	// Rule type: 0 (Internet border rule), 1 (inter-VPC rule), or 2 (NAT rule).
	Type int `json:"type,omitempty"`
	// Tag object attached to a rule.
	Tag TagsVO `json:"tag,omitempty"`
}

// This function is used to update an ACL rule.
func UpdateACLRule(client *golangsdk.ServiceClient, ruleId string, opts UpdateACLRuleOpts) (*RuleId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/acl-rule/{acl_rule_id}
	raw, err := client.Put(client.ServiceURL("acl-rule", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type UpdateResponse struct {
	// Rule Data
	Data RuleId `json:"data"`
}
