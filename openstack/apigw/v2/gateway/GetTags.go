package gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func GetTags(client *golangsdk.ServiceClient, instanceId string) ([]tags.ResourceTag, error) {
	var r struct {
		Tags []tags.ResourceTag `json:"tags"`
	}
	_, err := client.Get(client.ServiceURL("apigw", "instances", instanceId, "instance-tags"), &r, nil)
	return r.Tags, err
}
