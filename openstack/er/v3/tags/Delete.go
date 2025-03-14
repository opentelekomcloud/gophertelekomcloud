package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func Delete(client *golangsdk.ServiceClient, resourceType, resourceId string, key string) error {
	_, err := client.Delete(client.ServiceURL(resourceType, resourceId, "tags", key), &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
