package tags

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type tagsActionOpts struct {
	Action string             `json:"action" required:"true"`
	Tags   []tags.ResourceTag `json:"tags,omitempty"`
}

// Create an instance tags with given parameters.
func Create(client *golangsdk.ServiceClient, instanceID string, tags []tags.ResourceTag) error {
	opts := tagsActionOpts{
		Action: "create",
		Tags:   tags,
	}
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	url := client.ServiceURL("dcs", instanceID, "tags", "action")
	_, err = client.Post(strings.Replace(url, "v1.0", "v2", 1), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})
	return err
}
