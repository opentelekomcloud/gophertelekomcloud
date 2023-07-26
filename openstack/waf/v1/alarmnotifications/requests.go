package alarmnotifications

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToAlarmNotificationUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains all the values needed to update the alarm notification.
type UpdateOpts struct {
	Enabled       *bool    `json:"enabled" required:"true"`
	TopicURN      *string  `json:"topic_urn" required:"true"`
	SendFrequency int      `json:"sendfreq" required:"true"`
	Times         int      `json:"times" required:"true"`
	Threat        []string `json:"threat" required:"true"`
	Locale        string   `json:"locale,omitempty"`
}

// ToAlarmNotificationUpdateMap casts a UpdateOpts struct to a map.
func (opts UpdateOpts) ToAlarmNotificationUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Update is an operation which modifies the attributes of the specified alarm notification.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToAlarmNotificationUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(resourceURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})
	return
}

// List is method that can be able to list all alarm notification of WAF service
func List(client *golangsdk.ServiceClient) (r ListResult) {
	_, r.Err = client.Get(listURL(client), &r.Body, openstack.StdRequestOpts())
	return
}
