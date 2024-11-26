package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type ExpirationOptsBuilder interface {
	ToExpirationGetMap() (map[string]interface{}, error)
}

type ExpirationOpts struct {
	Duration int `json:"duration" required:"true"`
}

func (opts ExpirationOpts) ToExpirationGetMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// GetCert retrieves a particular cluster certificate based on its unique ID.
func GetCert(c *golangsdk.ServiceClient, id string) (r GetCertResult) {
	_, r.Err = c.Get(certificateURL(c, id), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}

// GetCertWithExpiration retrieves a particular cluster certificate based on its unique ID.
func GetCertWithExpiration(c *golangsdk.ServiceClient, id string, opts ExpirationOptsBuilder) (r GetCertResult) {
	b, err := opts.ToExpirationGetMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = c.Post(certificateURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}

type UpdateIpOpts struct {
	Action    string `json:"action" required:"true"`
	Spec      IpSpec `json:"spec,omitempty"`
	ElasticIp string `json:"elasticIp"`
}

type IpSpec struct {
	ID string `json:"id" required:"true"`
}

type UpdateIpOptsBuilder interface {
	ToMasterIpUpdateMap() (map[string]interface{}, error)
}

func (opts UpdateIpOpts) ToMasterIpUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "spec")
}

// Update the access information of a specified cluster.
func UpdateMasterIp(c *golangsdk.ServiceClient, id string, opts UpdateIpOptsBuilder) (r UpdateIpResult) {
	b, err := opts.ToMasterIpUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(masterIpURL(c, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
