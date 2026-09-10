package securitygrouprules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*SecurityGroupRule, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.ProjectID, "security-group-rules", id).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		SecurityGroupRule SecurityGroupRule `json:"security_group_rule"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.SecurityGroupRule, nil
}
