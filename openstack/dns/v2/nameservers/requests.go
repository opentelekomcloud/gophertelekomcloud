package nameservers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type ListNameserversRequest struct {
	Name string `json:"name,omitempty"`
}

func List(client *golangsdk.ServiceClient, zoneID string) (r GetResult) {
	url := baseURL(client, zoneID)
	_, r.Err = client.Get(url, &r.Body, nil)

	return
}
