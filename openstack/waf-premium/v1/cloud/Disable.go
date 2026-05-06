package cloud

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteOpts struct {
	// Enterprise project ID.
	EnterpriseProjectID string `json:"-" q:"enterprise_project_id,omitempty"`
}

func Disable(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("waf", "postpaid").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return err
	}

	// DELETE /v1/{project_id}/waf/postpaid
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json;charset=utf8",
			"region":       client.RegionID,
		},
	})
	return
}
