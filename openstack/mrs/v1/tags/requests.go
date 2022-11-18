package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func Get(client *golangsdk.ServiceClient, clusterID string) (r tags.GetResult) {
	return tags.Get(client, "clusters", clusterID)
}

func Create(client *golangsdk.ServiceClient, clusterID string, tag []tags.ResourceTag) (r tags.ActionResult) {
	return tags.Create(client, "clusters", clusterID, tag)
}

func Delete(client *golangsdk.ServiceClient, clusterID string, tag []tags.ResourceTag) (r tags.ActionResult) {
	return tags.Delete(client, "clusters", clusterID, tag)
}
