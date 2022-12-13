package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// OBS bucket name. The value contains 3 to 63 characters and must start with a digit or lowercase letter.
	// Only lowercase letters, digits, hyphens (-), and periods (.) are allowed.
	BucketName string `json:"bucket_name" required:"true"`
	// Prefix of trace files that need to be stored in OBS buckets. The value can contain 0 to 64 characters,
	// including letters, digits, hyphens (-), underscores (_), and periods (.).
	FilePrefixName string `json:"file_prefix_name,omitempty"`
	// Status of a tracker. The value can be enabled or disabled.
	// If you change the value to disabled, the tracker stops recording traces.
	Status string `json:"status,omitempty"`
	// Whether trace analysis is enabled.
	// When you enable trace analysis, a log group named CTS and a log stream named system-trace are created in LTS.
	Lts CreateLts `json:"lts,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts, trackerName string) (*Tracker, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("tracker", trackerName), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
