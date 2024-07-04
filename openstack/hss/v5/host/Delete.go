package hss

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteOpts struct {
	// Group ID
	GroupID string `q:"group_id"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (err error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("host-management", "groups").
		WithQueryParams(&opts).Build()
	if err != nil {
		return err
	}

	// DELETE /v5/{project_id}/host-management/groups
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"region": client.RegionID},
	})
	return
}
