package tag

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListClusterTags(client *golangsdk.ServiceClient, clusterId string) ([]TagPlain, error) {
	// GET /v1.0/{project_id}/clusters/{cluster_id}/tags
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "tags"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []TagPlain
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
