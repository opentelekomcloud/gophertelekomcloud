package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("premium-waf", "instance", id), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}
