package publicips

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*PublicIP, error) {
	raw, err := client.Get(client.ServiceURL(client.ProjectID, "publicips", id), nil, nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		PublicIP PublicIP `json:"publicip"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.PublicIP, err
}
