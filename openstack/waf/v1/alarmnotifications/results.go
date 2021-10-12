package alarmnotifications

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type AlarmNotification struct {
	ID            string   `json:"id"`
	Enabled       bool     `json:"enabled"`
	TopicURN      string   `json:"topic_urn"`
	SendFrequency int      `json:"sendfreq"`
	Times         int      `json:"times"`
	Threat        []string `json:"threat"`
	Locale        string   `json:"locale"`
}

type commonResult struct {
	golangsdk.Result
}

type ListResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

func (r commonResult) Extract() (*AlarmNotification, error) {
	s := new(AlarmNotification)
	err := r.ExtractIntoStructPtr(s, "")
	if err != nil {
		return nil, err
	}
	return s, nil
}
