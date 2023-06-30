package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

func ListClusterTags(client *golangsdk.ServiceClient, clusterId string) ([]tags.ResourceTag, error) {
	// GET /v1.0/{project_id}/clusters/{cluster_id}/tags
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []tags.ResourceTag
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
