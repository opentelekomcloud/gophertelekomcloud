package notifications

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// Delete removes a topic's notification configuration from an AS group.
func Delete(client *golangsdk.ServiceClient, scalingGroupID, topicURN string) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("scaling_notification", scalingGroupID, topicURN).Build()
	if err != nil {
		return err
	}
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{OkCodes: []int{204}})
	return err
}
