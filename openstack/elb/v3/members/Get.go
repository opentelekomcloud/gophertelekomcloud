package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, poolID, memberID string) (*Member, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("pools", poolID, "members", memberID).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res struct {
		Member Member `json:"member"`
	}
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res.Member, nil
}
