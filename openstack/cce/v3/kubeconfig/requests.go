package kubeconfig

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type GetOpts struct {
	Duration int `json:"duration" required:"true"`
}

func (opts GetOpts) ToCertificateGetMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(c *golangsdk.ServiceClient, clusterID string, opts GetOpts) (r GetResult) {
	b, err := opts.ToCertificateGetMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c, clusterID), b, &r.Body, openstack.StdRequestOpts())
	return
}
