package instances

import "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteOptsBuilder interface {
	ToInstanceDeleteQuery() (string, error)
}

type DeleteOpts struct {
	DeleteInstance bool `q:"instance_delete"`
}

func (opts DeleteOpts) ToInstanceDeleteQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func Delete(client *golangsdk.ServiceClient, id string, opts DeleteOptsBuilder) (r DeleteResult) {
	url := client.ServiceURL("scaling_group_instance", id)
	if opts != nil {
		q, err := opts.ToInstanceDeleteQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += q
	}
	_, r.Err = client.Delete(url, nil)
	return
}
