package members

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, poolID, memberID string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("pools", poolID, "members", memberID).Build()
	if err != nil {
		return err
	}
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{OkCodes: []int{204}})
	return err
}
