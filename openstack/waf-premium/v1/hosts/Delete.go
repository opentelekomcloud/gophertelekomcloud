package hosts

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteOpts struct {
	KeepPolicy *bool `q:"keepPolicy"`
}

func Delete(client *golangsdk.ServiceClient, id string, opts DeleteOpts) (err error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("premium-waf", "host", id).WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}
