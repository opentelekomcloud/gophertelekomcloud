package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateCcOpts struct {
	// Protection mode of the CC attack protection rule, which corresponds to the Mode
	// field in the Add CC Attack Protection Rule dialog box on the WAF console.
	// 0: standard. Only the protected paths of domain names can be specified.
	// 1: The path, IP address, cookie, header, and params fields can all be set.
	Mode *int `json:"mode" required:"true"`
	// Path to be protected in the CC attack protection rule.
	// This parameter is mandatory when the CC attack protection rule is in standard mode (i.e. the value of mode is 0).
	Url string `json:"url" required:"true"`
	// Rate limit conditions of the CC protection rule.
	// This parameter is mandatory when the CC protection rule is in advanced mode (i.e. the value of mode is 1).
	Conditions []CcConditionsObject `json:"conditions"`
	// Protection action to take if the number of requests reaches the upper limit.
	Action *CcActionObject `json:"action" required:"true"`
	// Rate limit mode.
	// ip: IP-based rate limiting. Website visitors are identified by IP address.
	// cookie: User-based rate limiting. Website visitors are identified by the cookie key value.
	// header: User-based rate limiting. Website visitors are identified by the header field.
	// other: Website visitors are identified by the Referer field (user-defined request source).
	TagType string `json:"tag_type" required:"true"`
	// User identifier. This parameter is mandatory when the rate limit mode is set to user (cookie or header).
	// cookie: Set the cookie field name.
	// You need to configure an attribute variable name in the cookie that can uniquely identify
	// a web visitor based on your website requirements. This field does not support regular expressions.
	// Only complete matches are supported. For example, if a website uses the name field
	// in the cookie to uniquely identify a website visitor, select name.
	// header: Set the user-defined HTTP header you want to protect.
	// You need to configure the HTTP header that can identify web visitors based on your website requirements.
	TagIndex string `json:"tag_index"`
	// User tag. This parameter is mandatory when the rate limit mode is set to other.
	// other: A website visitor is identified by the Referer field (user-defined request source).
	TagCondition *CcTagConditionObject `json:"tag_condition"`
	// Rate limit frequency based on the number of requests. The value ranges from 1 to 2,147,483,647.
	LimitNum int64 `json:"limit_num" required:"true"`
	// Rate limit period, in seconds. The value ranges from 1 to 3,600.
	LimitPeriod int64 `json:"limit_period" required:"true"`
	// Allowable frequency based on the number of requests. The value ranges from 0 to 2,147,483,647.
	// This parameter is required only when the protection action type is dynamic_block.
	UnlockNum int64 `json:"unlock_num"`
	// Block duration, in seconds. The value ranges from 0 to 65,535.
	// Specifies the period within which access is blocked. An error page is displayed in this period.
	LockTime int `json:"lock_time"`
	// Rule description.
	Description string `json:"description"`
}

type CcConditionsObject struct {
	// Field type. The value can be url, ip, params, cookie, or header.
	Category string `json:"category" required:"true"`
	// Logic for matching the condition.
	// If the category is url, the optional operations are:
	// `contain`, `not_contain`, `equal`, `not_equal`, `prefix`, `not_prefix`, `suffix`, `not_suffix`,
	// `contain_any`, `not_contain_all`, `equal_any`, `not_equal_all`, `equal_any`,
	// `not_equal_all`, `prefix_any`, `not_prefix_all`, `suffix_any`, `not_suffix_all`,
	// `len_greater`, `len_less`, `len_equal` and `len_not_equal`
	// If the category is ip, the optional operations are:
	// `equal`, `not_equal`, `equal_any` and `not_equal_all`
	// If the category is params, cookie and header, the optional operations are:
	// `contain`, `not_contain`, `equal`, `not_equal`, `prefix`, `not_prefix`, `suffix`, `not_suffix`,
	// `contain_any`, `not_contain_all`, `equal_any`, `not_equal_all`, `equal_any`, `not_equal_all`,
	// `prefix_any`, `not_prefix_all`, `suffix_any`, `not_suffix_all`, `len_greater`, `len_less`,
	// `len_equal`, `len_not_equal`, `num_greater`, `num_less`, `num_equal`, `num_not_equal`,
	// `exist` and `not_exist`
	LogicOperation string `json:"logic_operation" required:"true"`
	// Content of the conditions.
	// This parameter is mandatory when the suffix of logic_operation is not any or all.
	Contents []string `json:"contents"`
	// Reference table ID. It can be obtained by calling the API Querying the Reference Table List.
	// This parameter is mandatory when the suffix of logic_operation is any or all.
	// The reference table type must be the same as the category type.
	ValueListId string `json:"value_list_id"`
	// Subfield. When category is set to params, cookie, or header,
	// set this parameter based on site requirements.
	// This parameter is mandatory.
	Index string `json:"index"`
}

type CcActionObject struct {
	// Action type:
	// captcha: Verification code.
	// WAF requires visitors to enter a correct verification code to continue their
	// access to requested page on your website.
	// block: WAF blocks the requests. When tag_type is set to other, the value can only be block.
	// log: WAF logs the event only.
	// dynamic_block: In the previous rate limit period,
	// if the request frequency exceeds the value of Rate Limit Frequency,
	// the request is blocked. In the next rate limit period,
	// if the request frequency exceeds the value of Permit Frequency,
	// the request is still blocked.
	// Note: The dynamic_block protection action can be set only when the
	// advanced protection mode is enabled for the CC protection rule.
	Category string `json:"category" required:"true"`
	// Block page information. When protection action category is set to block or dynamic_block,
	// you need to set the returned block page.
	// If you want to use the default block page, this parameter can be excluded.
	// If you want to use a custom block page, set this parameter.
	Detail *CcDetailObject `json:"detail"`
}

type CcDetailObject struct {
	// Returned page.
	Response *CcResponseObject `json:"response"`
}

type CcResponseObject struct {
	// Content type. The value can only be application/json, text/html, or text/xml.
	ContentType string `json:"content_type"`
	// Protection page content.
	Content string `json:"content"`
}

type CcTagConditionObject struct {
	// User identifier. The value is fixed at referer.
	Category string `json:"category"`
	// Content of the user identifier field.
	Contents []string `json:"contents"`
}

// CreateCc will create a cc rule on the values in CreateOpts.
func CreateCc(client *golangsdk.ServiceClient, policyId string, opts CreateCcOpts) (*CcRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/cc
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "cc"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res CcRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CcRule struct {
	// Rule ID.
	ID string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// When the value of mode is 0, this parameter has a return value.
	Url string `json:"url"`
	// Whether a prefix is used for the path.
	// If the protected URL ends with an asterisk (*), a path prefix is used.
	Prefix bool `json:"prefix"`
	// Mode.
	// 0: Standard.
	// 1: Advanced.
	Mode int `json:"mode"`
	// Rule status. The value can be 0 or 1.
	// 0: The rule is disabled.
	// 1: The rule is enabled.
	Status int `json:"status"`
	// Rate limit conditions of the CC protection rule.
	Conditions []CcConditionsObject `json:"conditions"`
	// Protection action to take if the number of requests reaches the upper limit.
	Action CcActionObject `json:"action"`
	// Rate limit mode.
	TagType string `json:"tag_type"`
	// User identifier.
	// This parameter is mandatory when the rate limit mode is set to user (cookie or header).
	TagIndex string `json:"tag_index"`
	// User tag.
	TagCondition CcTagConditionObject `json:"tag_condition"`
	// Rate limit frequency based on the number of requests. The value ranges from 1 to 2,147,483,647.
	LimitNum int64 `json:"limit_num"`
	// Rate limit period, in seconds. The value ranges from 1 to 3,600.
	LimitPeriod int `json:"limit_period"`
	// Allowable frequency based on the number of requests.
	UnlockNum int64 `json:"unlock_num"`
	// Block duration, in seconds.
	LockTime int64 `json:"lock_time"`
	// Rule description.
	Description string `json:"description"`
	// This parameter is reserved and can be ignored currently.
	TotalNum int `json:"total_num"`
	// This parameter is reserved and can be ignored currently.
	UnAggregation bool `json:"unaggregation"`
	// Rule aging time. This parameter is reserved and can be ignored currently.
	AgingTime int `json:"aging_time"`
	// Rule creation object. This parameter is reserved and can be ignored currently.
	Producer int `json:"producer"`
	// Timestamp the rule is created.
	CreatedAt int64 `json:"timestamp"`
}
