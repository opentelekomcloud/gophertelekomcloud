package subscriptions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// CreateOptsBuilder is used for creating subscription parameters.
// any struct providing the parameters should implement this interface
type CreateOptsBuilder interface {
	ToSubscriptionCreateMap() (map[string]interface{}, error)
}

// CreateOpts is a struct that contains all the parameters.
type CreateOpts struct {
	// Message endpoint
	Endpoint string `json:"endpoint" required:"true"`
	// Protocol of the message endpoint
	Protocol string `json:"protocol" required:"true"`
	// Description of the subscription
	Remark string `json:"remark,omitempty"`
}

func (ops CreateOpts) ToSubscriptionCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(ops, "")
}

// Create a subscription with given parameters.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder, topicUrn string) (r CreateResult) {
	b, err := opts.ToSubscriptionCreateMap()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = client.Post(createURL(client, topicUrn), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{201, 200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})

	return
}

// Delete a subscription via subscription urn
func Delete(client *golangsdk.ServiceClient, subscriptionUrn string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, subscriptionUrn), &golangsdk.RequestOpts{
		OkCodes:     []int{200, 202, 204},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})
	return
}

// List all the subscriptions
func List(client *golangsdk.ServiceClient) (r ListResult) {
	_, r.Err = client.Get(listURL(client), &r.Body, openstack.StdRequestOpts())
	return
}

// ListFromTopic all the subscriptions from topic
func ListFromTopic(client *golangsdk.ServiceClient, subscriptionUrn string) (r ListResult) {
	_, r.Err = client.Get(listFromTopicURL(client, subscriptionUrn), &r.Body, openstack.StdRequestOpts())
	return
}
