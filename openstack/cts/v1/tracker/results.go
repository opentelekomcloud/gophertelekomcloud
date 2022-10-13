package tracker

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Tracker struct {
	// Status of a tracker. The value is enabled.
	Status string `json:"status"`
	// OBS bucket name. The value contains 3 to 63 characters and must start with a digit or lowercase letter.
	// Only lowercase letters, digits, hyphens (-), and periods (.) are allowed.
	BucketName string `json:"bucket_name"`
	// Prefix of trace files that need to be stored in OBS buckets.
	FilePrefixName string `json:"file_prefix_name"`
	// Tracker name. The default value is system.
	TrackerName string `json:"tracker_name"`
	Lts         Lts    `json:"lts"`

	IsSupportTraceFilesEncryption bool   `json:"is_support_trace_files_encryption"`
	CreateTime                    int64  `json:"create_time"`
	StreamId                      string `json:"streamId"`
	GroupId                       string `json:"groupId"`
	IsSupportValidate             bool   `json:"is_support_validate"`
	TrackerType                   string `json:"tracker_type"`
	DomainId                      string `json:"domain_id"`
	ProjectId                     string `json:"project_id"`
	Id                            string `json:"id"`
}

type Lts struct {
	// Whether trace analysis is enabled.
	IsLtsEnabled bool `json:"is_lts_enabled"`
	// Name of the Log Tank Service (LTS) log group.
	LogGroupName string `json:"log_group_name"`
	// Name of the LTS log stream.
	LogTopicName string `json:"log_topic_name"`
}

func extra(err error, raw *http.Response) (*Tracker, error) {
	if err != nil {
		return nil, err
	}

	var res Tracker
	err = extract.Into(raw.Body, &res)
	return &res, err
}
