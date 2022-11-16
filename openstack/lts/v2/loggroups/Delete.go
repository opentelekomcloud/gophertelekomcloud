package loggroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete a log group by id
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	opts := golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		OkCodes:     []int{204},
	}
	_, r.Err = client.Delete(client.ServiceURL("log-groups", id), &opts)
	return
}
