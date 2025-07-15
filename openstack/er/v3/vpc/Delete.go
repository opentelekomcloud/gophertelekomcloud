package vpc

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, erId, vpcId string) (err error) {
	_, err = client.Delete(client.ServiceURL("enterprise-router", erId, "vpc-attachments", vpcId), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
