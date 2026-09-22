package pools

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("pools", id).Build()
	if err != nil {
		return err
	}
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{OkCodes: []int{204}})
	return err
}
