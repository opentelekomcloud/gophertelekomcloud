package recorder

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Recorder domain id
	DomainId string `json:"-"`
	// Specifies configurations for the tracker channel.
	Channel ChannelConfigBody `json:"channel" required:"true"`
	// Specifies the selector.
	Selector SelectorConfigBody `json:"selector" required:"true"`
	// Number of days for data storage.
	RetentionDays *int `json:"retention_period_in_days,omitempty"`
	// Specifies the IAM agency name.
	AgencyName string `json:"agency_name" required:"true"`
}

func UpdateRecorder(client *golangsdk.ServiceClient, opts UpdateOpts) error {
	// PUT /v1/resource-manager/domains/{domain_id}/tracker-config
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("resource-manager", "domains", opts.DomainId, "tracker-config"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
