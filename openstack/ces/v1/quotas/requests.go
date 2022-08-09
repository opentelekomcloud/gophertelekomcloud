package quotas

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func ShowQuotas(client *golangsdk.ServiceClient) (r ShowQuotasResult) {
	_, r.Err = client.Get(quotasURL(client), &r.Body, nil)
	return
}
