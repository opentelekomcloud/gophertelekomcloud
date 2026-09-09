package securitygrouprules

import (
	"bytes"
	"fmt"
	"strconv"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

const defaultPageLimit = 2000

type ListOpts struct {
	Marker          string `q:"marker,omitempty"`
	Limit           *int   `q:"limit,omitempty"`
	SecurityGroupID string `q:"security_group_id,omitempty"`
	RemoteIPPrefix  string `q:"remote_ip_prefix,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]SecurityGroupRule, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "security-group-rules").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SecurityGroupRulePage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractSecurityGroupRules(pages)
}

type SecurityGroupRulePage struct {
	pagination.NewPageResult
}

func (p SecurityGroupRulePage) NewNextPageURL() (string, error) {
	rules, err := ExtractSecurityGroupRules(p)
	if err != nil || len(rules) == 0 {
		return "", err
	}
	limit := defaultPageLimit
	if value := p.URL.Query().Get("limit"); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil {
			return "", fmt.Errorf("invalid security group rule page limit %q: %w", value, err)
		}
	}
	if limit > 0 && len(rules) < limit {
		return "", nil
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", rules[len(rules)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p SecurityGroupRulePage) NewIsEmpty() (bool, error) {
	rules, err := ExtractSecurityGroupRules(p)
	return len(rules) == 0, err
}

func ExtractSecurityGroupRules(page pagination.NewPage) ([]SecurityGroupRule, error) {
	var res struct {
		SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
	}
	err := extract.Into(bytes.NewReader(page.(SecurityGroupRulePage).Body), &res)
	return res.SecurityGroupRules, err
}
