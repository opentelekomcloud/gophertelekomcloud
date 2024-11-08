package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func Delete(client *golangsdk.ServiceClient, resourceType, resourceId string, opts TagsOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	_, err = client.Post(client.ServiceURL(resourceType, resourceId, "tags", "delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	return err
}
