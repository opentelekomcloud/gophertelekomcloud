package transfers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateLogDumpObsOpts struct {
	// Log group ID.
	// Value length: 36 characters
	LogGroupId string `json:"log_group_id" required:"true"`
	// Indicates IDs of log streams whose logs are to be periodically transferred to OBS. You can specify one or more log streams.
	// Example value:
	// 7bb6b1e7-xxxx-4255-87f9-b3dc7fb2xxxx
	LogStreamIds []string `json:"log_stream_ids" required:"true"`
	// Indicates the name of an OBS bucket.
	// Minimum length: 3 characters
	// Maximum length: 63 characters
	ObsBucketName string `json:"obs_bucket_name" required:"true"`
	// Set this parameter to cycle, which indicates that the log transfer is periodic.
	// Value length: 5 characters
	Type string `json:"type" required:"true"`
	// Indicates whether the logs are stored in raw or JSON format. The default value is RAW.
	// Minimum length: 3 characters
	// Maximum length: 4 characters
	StorageFormat string `json:"storage_format" required:"true"`
	// Indicates whether the log transfer is enabled. The value is true (default) or false.
	SwitchOn *bool `json:"switch_on,omitempty"`
	// Indicates the file name prefix of the log files transferred to an OBS bucket.
	// Minimum length: 0 characters
	// Maximum length: 64 characters
	PrefixName string `json:"prefix_name,omitempty"`
	// Indicates a custom path to store the log files.
	// Minimum length: 0 characters
	// Maximum length: 64 characters
	DirPrefixName string `json:"dir_prefix_name,omitempty"`
	// Indicates the length of the log transfer interval.
	// Example values: 1, 2, 3, 5, 6, 12, and 30
	Period int32 `json:"period" required:"true"`
	// Indicates the unit of the log transfer interval.
	// Example values: min and hour
	// Minimum length: 3 characters
	// Maximum length: 4 characters
	// NOTE:
	// The log transfer interval is specified by the combination of the values of period and period_unit, and must be set to one of the following: 2 min, 5 min, 30 min, 1 hour, 3 hours, 6 hours, and 12 hours.
	PeriodUnit string `json:"period_unit" required:"true"`
}

func CreateLogDumpObs(client *golangsdk.ServiceClient, opts CreateLogDumpObsOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/log-dump/obs
	raw, err := client.Post(client.ServiceURL("log-dump", "obs"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"log_dump_obs_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
