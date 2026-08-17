package favorites

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteOpts struct {
	// ResourceID is the ID of the favorite resource to remove.
	ResourceID string
}

func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) error {
	if opts.ResourceID == "" {
		return fmt.Errorf("favorite resource ID must not be empty")
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("lts", "favorite", opts.ResourceID).Build()
	if err != nil {
		return err
	}

	// DELETE /v1.0/{project_id}/lts/favorite/{fav_res_id}
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json;charset=UTF-8",
		},
		OkCodes: []int{200},
	})
	return err
}
