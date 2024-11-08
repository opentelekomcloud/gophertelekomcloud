package recorder

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetRecorder(client *golangsdk.ServiceClient, domainId string) (*Recorder, error) {
	// GET /v1/resource-manager/domains/{domain_id}/tracker-config
	raw, err := client.Get(client.ServiceURL(
		"resource-manager", "domains", domainId,
		"tracker-config"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Recorder
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Recorder struct {
	Channel         ChannelConfigBody  `json:"channel"`
	Selector        SelectorConfigBody `json:"selector"`
	RetentionPeriod int                `json:"retention_period_in_days"`
	AgencyName      string             `json:"agency_name"`
}

type ChannelConfigBody struct {
	Smn *TrackerSMNConfigBody `json:"smn,omitempty"`
	Obs *TrackerObsConfigBody `json:"obs,omitempty"`
}

type TrackerSMNConfigBody struct {
	RegionId  string `json:"region_id"`
	ProjectId string `json:"project_id"`
	TopicUrn  string `json:"topic_urn"`
}

type TrackerObsConfigBody struct {
	BucketName   string  `json:"bucket_name"`
	BucketPrefix *string `json:"bucket_prefix,omitempty"`
	RegionId     string  `json:"region_id"`
}

type SelectorConfigBody struct {
	AllSupported  bool     `json:"all_supported"`
	ResourceTypes []string `json:"resource_types"`
}
