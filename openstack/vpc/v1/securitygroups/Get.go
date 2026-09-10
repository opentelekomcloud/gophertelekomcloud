package securitygroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*SecurityGroup, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.ProjectID, "security-groups", id).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		SecurityGroup SecurityGroup `json:"security_group"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.SecurityGroup, nil
}
