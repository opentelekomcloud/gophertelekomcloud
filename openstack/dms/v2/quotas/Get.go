package quotas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type GetOpts struct {
	// Whether to include the tag quota. Defaults to true.
	IncludeTagsQuota *bool `q:"includeTagsQuota,omitempty"`
	// Engine to query the quota of: reliability (RocketMQ) or kafka.
	OnlyQuota string `q:"onlyQuota,omitempty"`
}

type Quotas struct {
	// Quota list.
	Resources []Resource `json:"resources"`
}

type Resource struct {
	// Quota type.
	//    kafkaInstance: Kafka instance quotas.
	//    rocketmqInstance: RocketMQ instance quotas.
	//    tags: tag quotas.
	Type string `json:"type"`
	// Maximum number of instances a tenant can create,
	// or maximum number of tags of each instance.
	Quota int `json:"quota"`
	// Number of created instances. Not returned for tags.
	Used int `json:"used"`
}

// Get the DMS instance and tag quotas of the tenant.
// Send GET to /v2/{project_id}/quotas
func Get(client *golangsdk.ServiceClient, opts GetOpts) (*Quotas, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("quotas").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Quotas
	err = extract.IntoStructPtr(raw.Body, &res, "quotas")
	return &res, err
}
