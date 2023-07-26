package nodes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// List returns collection of nodes.
func List(client *golangsdk.ServiceClient, clusterID string) (r ListResult) {
	_, r.Err = client.Get(listURL(client, clusterID), &r.Body, openstack.StdRequestOpts())
	return
}

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(c *golangsdk.ServiceClient, clusterID, k8sName string) (r GetResult) {
	_, r.Err = c.Get(nodeURL(c, clusterID, k8sName), &r.Body, openstack.StdRequestOpts())
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToNodeUpdateMap() (map[string]any, error)
}

// UpdateOpts contains all the values needed to update a new node
type UpdateOpts struct {
	Metadata Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	Labels map[string]any `json:"labels,omitempty"`
}

// ToNodeUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToNodeUpdateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update allows nodes to be updated.
func Update(c *golangsdk.ServiceClient, clusterID, k8sName string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToNodeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Patch(nodeURL(c, clusterID, k8sName), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/merge-patch+json",
		},
	})
	return
}
