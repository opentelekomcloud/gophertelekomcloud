package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type StatusOpts struct {
	// An indicator of whether the number of DCS instances that failed to be created will be returned to the API caller. Options:
	// true: The number of DCS instances that failed to be created will be returned to the API caller.
	// false or others: The number of DCS instances that failed to be created will not be returned to the API caller.
	IncludeFailure *bool `q:"includeFailure"`
}

func ListDCSStatus(client *golangsdk.ServiceClient, opts StatusOpts) (*ListNumberOfInstancesInDifferentStatusResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", "status").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListNumberOfInstancesInDifferentStatusResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListNumberOfInstancesInDifferentStatusResponse struct {
	// Number of instances for which payment is in progress.
	PayingCount int32 `json:"paying_count"`
	// Number of instances for which payment is in progress.
	MigratingCount int32 `json:"migrating_count"`
	// Number of instances whose data is being cleared.
	FlushingCount int32 `json:"flushing_count"`
	// Number of instances that are being upgraded.
	UpgradingCount int32 `json:"upgrading_count"`
	// Number of instances for which data restoration is in progress.
	RestoringCount int32 `json:"restoring_count"`
	// Number of instances that are being scaled up.
	ExtendingCount int32 `json:"extending_count"`
	// Number of instances that are being created.
	CreatingCount int32 `json:"creating_count"`
	// Number of running instances.
	RunningCount int32 `json:"running_count"`
	// Number of abnormal instances.
	ErrorCount int32 `json:"error_count"`
	// Number of instances that fail to be created.
	CreatefailedCount int32 `json:"createfailed_count"`
	// Number of instances that are being restarted.
	RestartingCount int32 `json:"restarting_count"`
	// Number of instances that are being deleted.
	DeletingCount int32 `json:"deleting_count"`
	// Number of instances that have been stopped.
	ClosedCount int32 `json:"closed_count"`
	// Number of instances that are being started.
	StartingCount int32 `json:"starting_count"`
	// Number of instances that are being stopped.
	ClosingCount int32 `json:"closing_count"`
}
