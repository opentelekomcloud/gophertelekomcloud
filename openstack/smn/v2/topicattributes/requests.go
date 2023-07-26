package topicattributes

import "github.com/opentelekomcloud/gophertelekomcloud"

func commonOpts() *golangsdk.RequestOpts {
	return &golangsdk.RequestOpts{
		OkCodes: []int{200},
	}
}

type ListOptsBuilder interface {
	ToAttributeListQuery() (string, error)
}

type ListOpts struct {
	Name string `q:"name"`
}

func (opts ListOpts) ToAttributeListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

func List(client *golangsdk.ServiceClient, topicURN string, opts ListOptsBuilder) (r GetResult) {
	url := listURL(client, topicURN)
	if opts != nil {
		q, err := opts.ToAttributeListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += q
	}
	_, r.Err = client.Get(url, &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToAttributeUpdateMap() (map[string]any, error)
}

type UpdateOpts struct {
	Value string `json:"value"`
}

func (opts UpdateOpts) ToAttributeUpdateMap() (map[string]any, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

func Update(client *golangsdk.ServiceClient, topicURN, attribute string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToAttributeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(attributeURL(client, topicURN, attribute), b, &r.Body, commonOpts())
	return
}

func Delete(client *golangsdk.ServiceClient, topicURN, attribute string) (r DeleteResult) {
	_, r.Err = client.Delete(attributeURL(client, topicURN, attribute), commonOpts())
	return
}
