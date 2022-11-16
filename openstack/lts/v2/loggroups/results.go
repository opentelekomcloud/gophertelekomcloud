package loggroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Log group Create response

// DeleteResult is a struct which contains the result of deletion
type DeleteResult struct {
	golangsdk.ErrResult
}

// LogGroup response
type LogGroup struct {
	ID           string `json:"log_group_id"`
	Name         string `json:"log_group_name"`
	CreationTime int64  `json:"creation_time"`
	TTLinDays    int    `json:"ttl_in_days"`
}

// GetResult contains the body of getting detailed
type GetResult struct {
	golangsdk.Result
}

// Extract from GetResult
func (r GetResult) Extract() (*LogGroup, error) {
	s := new(LogGroup)
	err := r.Result.ExtractInto(s)
	return s, err
}
