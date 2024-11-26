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

// Delete will permanently delete a particular cluster based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	return
}

type DeleteOpts struct {
	ErrorStatus string `q:"errorStatus"`
	DeleteEfs   string `q:"delete_efs"`
	DeleteENI   string `q:"delete_eni"`
	DeleteEvs   string `q:"delete_evs"`
	DeleteNet   string `q:"delete_net"`
	DeleteObs   string `q:"delete_obs"`
	DeleteSfs   string `q:"delete_sfs"`
}

func DeleteWithOpts(c *golangsdk.ServiceClient, id string, opts DeleteOpts) error {
	url := resourceURL(c, id)
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return err
	}

	_, err = c.Delete(url+q.String(), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	return err
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
