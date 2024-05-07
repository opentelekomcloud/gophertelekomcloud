package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const statusEndpoint = "status"

// GetStatus is used to query the job status.
// Send request GET /v1.1/{project_id}/clusters/{cluster_id}/cdm/job/{job_name}/status
func GetStatus(client *golangsdk.ServiceClient, clusterId, jobName string, opts *GetOpts) (*StatusResp, error) {
	raw, err := client.Get(
		client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, jobEndpoint, jobName, statusEndpoint),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var res StatusResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type StatusResp struct {
	// Submissions is an array of StartJobSubmission objects.
	Submissions []*StatusJobSubmission `json:"submissions"`
}

type StatusJobSubmission struct {
	// IsIncrementing indicates whether the job migrates incremental data.
	IsIncrementing bool `json:"isIncrementing"`

	// JobName is the name of the job.
	JobName string `json:"job-name"`

	// Counters is Job's running result statistics. This parameter is available only when status is SUCCEEDED. For details, see the description of the counters parameter.
	Counters *Counters `json:"counters"`

	// IsStopingIncrement indicates whether to stop incremental data migration.
	IsStopingIncrement string `json:"isStopingIncrement"`

	// IsExecuteAuto indicates whether to execute the job as scheduled.
	IsExecuteAuto bool `json:"is-execute-auto"`

	// LastUpdateDate is the time when the job was last updated.
	LastUpdateDate int64 `json:"last-update-date"`

	// LastUpdateUser is the user who last updated the job status.
	LastUpdateUser string `json:"last-udpate-user"`

	// IsDeleteJob indicates whether to delete the job after it is executed.
	IsDeleteJob bool `json:"isDeleteJob"`

	// CreationUser is the user who created the job.
	CreationUser string `json:"creation-user"`

	// CreationDate is the time when the job was created, accurate to millisecond.
	CreationDate int64 `json:"creation-date"`

	// ExternalID is an external ID associated with the job (optional).
	ExternalID string `json:"external-id,omitempty"`

	// Progress is the job progress. If a job fails, the value is -1. Otherwise, the value ranges from 0 to 100.
	Progress float64 `json:"progress"`

	// SubmissionID is the ID of the submitted job.
	SubmissionID int `json:"submission-id"`

	// DeleteRows is the number of rows deleted during the job.
	DeleteRows int `json:"delete_rows"`

	// UpdateRows is the number of rows updated during the job.
	UpdateRows int `json:"update_rows"`

	// WriteRows is the number of rows written during the job.
	WriteRows int `json:"write_rows"`

	// ExecuteDate is the time when the job was executed.
	ExecuteDate int64 `json:"execute-date"`

	// Status represents the job's execution status.
	Status string `json:"status"`

	// ErrorDetails represents Error details. This parameter is available only when status is FAILED.
	ErrorDetails string `json:"error-details"`

	// ErrorSummary represents Error summary. This parameter is available only when status is FAILED.
	ErrorSummary string `json:"error-summary"`
}

type Counters struct {
	// BytesWritten is number of bytes that are written.
	BytesWritten int64 `json:"BYTES_WRITTEN"`
	// TotalFiles is the total number of files.
	TotalFiles int `json:"TOTAL_FILES"`
	// BytesWritten is number of rows that are read.
	RowsRead int64 `json:"ROWS_READ"`
	// BytesRead is number of bytes that are read
	BytesRead int64 `json:"BYTES_READ"`
	// RowsWritten is number of rows that are written.
	RowsWritten int64 `json:"ROWS_WRITTEN"`
	// FilesWritten is number of files that are written.
	FilesWritten int `json:"FILES_WRITTEN"`
	// FilesRead is number of files that are read.
	FilesRead int `json:"FILES_READ"`
	// TotalSize is total number of bytes.
	TotalSize int64 `json:"TOTAL_SIZE"`
	// FilesSkipped is number of files that are skipped.
	FilesSkipped int `json:"FILES_SKIPPED"`
	// RowsWrittenSkipped is number of rows that are skipped.
	RowsWrittenSkipped int64 `json:"ROWS_WRITTEN_SKIPPED"`
}
