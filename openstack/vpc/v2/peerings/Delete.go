package peerings

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, peeringID string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("vpc", "peerings", peeringID).Build()
	if err != nil {
		return err
	}
	_, err = client.Delete(client.ServiceURL(url.String()), nil)
	return err
}
