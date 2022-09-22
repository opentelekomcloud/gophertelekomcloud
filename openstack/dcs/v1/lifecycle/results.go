package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// InstanceCreate response
type InstanceCreate struct {
	InstanceID string `json:"instance_id"`
}

type ListDcsResponse struct {
	// Array of DCS instance details.
	Instances []Instance `json:"instances"`
	// Number of DCS instances.
	TotalCount int `json:"instance_num"`
}

// Password response
type Password struct {
	// Whether the password is successfully changed:
	// Values:
	// Success: The password is successfully changed.
	// passwordFailed: The old password is incorrect.
	// Locked: This account has been locked.
	// Failed: Failed to change the password.
	Result         string `json:"result"`
	Message        string `json:"message"`
	RetryTimesLeft string `json:"retry_times_left"`
	LockTime       string `json:"lock_time"`
	LockTimesLeft  string `json:"lock_time_left"`
}

// UpdatePasswordResult is a struct from which can get the result of update password method
type UpdatePasswordResult struct {
	golangsdk.Result
}

// Extract from UpdatePasswordResult
func (r UpdatePasswordResult) Extract() (*Password, error) {
	var s Password
	err := r.Result.ExtractInto(&s)
	return &s, err
}
