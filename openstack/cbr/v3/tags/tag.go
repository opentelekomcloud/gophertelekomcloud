package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func ShowVaultProjectTag(client *golangsdk.ServiceClient) (r tags.ListResult) {
	return tags.List(client, "vault")
}

func ShowVaultTag(client *golangsdk.ServiceClient, id string) (r tags.GetResult) {
	return tags.Get(client, "vault", id)
}

func CreateVaultTags(client *golangsdk.ServiceClient, id string, req []tags.ResourceTag) (r tags.ActionResult) {
	return tags.Create(client, "vault", id, req)
}

func DeleteVaultTag(client *golangsdk.ServiceClient, id string, req []tags.ResourceTag) (r tags.ActionResult) {
	return tags.Delete(client, "vault", id, req)
}
