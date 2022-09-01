package nodes

import (
	"fmt"
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToNodeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update a new node
type UpdateOpts struct {
	Metadata Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	Labels map[string]interface{} `json:"labels,omitempty"`
}

// ToNodeUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToNodeUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update allows nodes to be updated.
func Update(client *golangsdk.ServiceClient, clusterID, k8sName string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToNodeUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Patch(
		fmt.Sprintf("https://%s.%s", clusterID, client.ResourceBaseURL()[8:])+strings.Join([]string{"nodes", k8sName}, "/"),
		b, nil, &golangsdk.RequestOpts{
			OkCodes: []int{200},
			MoreHeaders: map[string]string{
				"Content-Type": "application/merge-patch+json",
			},
		})
	return
}
