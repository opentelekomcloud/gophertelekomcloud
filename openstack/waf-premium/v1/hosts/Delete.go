package hosts

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteOpts struct {
	KeepPolicy *bool `q:"keepPolicy"`
}

func Delete(client *golangsdk.ServiceClient, id string, opts DeleteOpts) (err error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return
	}
	_, err = client.Delete(client.ServiceURL("premium-waf", "host", id)+q.String(), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}
