package privateips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*PrivateIP, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "privateips", id).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		PrivateIP PrivateIP `json:"privateip"`
	}
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}
	return &res.PrivateIP, err
}
