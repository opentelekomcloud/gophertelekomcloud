package nodes

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains all the values needed to update a new node
type UpdateOpts struct {
	Metadata Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	Labels map[string]interface{} `json:"labels,omitempty"`
}

// Update allows nodes to be updated.
func Update(client *golangsdk.ServiceClient, clusterID, k8sName string, opts UpdateOpts) (*GetNode, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Patch(
		fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+
			strings.Join([]string{"nodes", k8sName}, "/"), b, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
			MoreHeaders: map[string]string{
				"Content-Type": "application/merge-patch+json",
			},
		})
	if err != nil {
		return nil, err
	}

	var res GetNode
	err = extract.Into(raw.Body, &res)
	return &res, err
}
