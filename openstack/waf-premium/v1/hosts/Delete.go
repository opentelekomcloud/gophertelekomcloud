package hosts

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	KeepPolicy *bool `q:"keepPolicy"`
}

func Delete(client *golangsdk.ServiceClient, id string, opts DeleteOpts) (err error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return
	}
	_, err = client.Delete(client.ServiceURL("premium-waf", "host", id)+q.String(), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}
