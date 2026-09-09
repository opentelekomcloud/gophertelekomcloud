package securitygroups

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
	Marker              string `q:"marker,omitempty"`
	Limit               *int   `q:"limit,omitempty"`
	VpcID               string `q:"vpc_id,omitempty"`
	EnterpriseProjectID string `q:"enterprise_project_id,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]SecurityGroup, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "security-groups").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SecurityGroupPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractSecurityGroups(pages)
}

type SecurityGroupPage struct {
	pagination.NewPageResult
}

func (p SecurityGroupPage) NewNextPageURL() (string, error) {
	groups, err := ExtractSecurityGroups(p)
	if err != nil || len(groups) == 0 {
		return "", err
	}
	limit := defaultPageLimit
	if value := p.URL.Query().Get("limit"); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil {
			return "", fmt.Errorf("invalid security group page limit %q: %w", value, err)
		}
	}
	if limit > 0 && len(groups) < limit {
		return "", nil
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", groups[len(groups)-1].ID)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p SecurityGroupPage) NewIsEmpty() (bool, error) {
	groups, err := ExtractSecurityGroups(p)
	return len(groups) == 0, err
}

func ExtractSecurityGroups(page pagination.NewPage) ([]SecurityGroup, error) {
	var res struct {
		SecurityGroups []SecurityGroup `json:"security_groups"`
	}
	err := extract.Into(bytes.NewReader(page.(SecurityGroupPage).Body), &res)
	return res.SecurityGroups, err
}
