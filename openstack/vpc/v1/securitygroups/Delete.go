package securitygroups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, id string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(client.ProjectID, "security-groups", id).Build()
	if err != nil {
		return err
	}
	_, err = client.Delete(client.ServiceURL(url.String()), nil)
	return err
}
