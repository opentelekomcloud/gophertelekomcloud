package transfers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateTransferOpts struct {
	TransferId   string       `json:"log_transfer_id"`
	TransferInfo TransferInfo `json:"log_transfer_info"`
}

type TransferInfo struct {
	// Log transfer format. The value can be RAW or JSON. RAW indicates raw log format, whereas JSON indicates JSON format. JSON and RAW are supported for OBS and DIS transfer tasks, but only RAW is supported for DMS transfer tasks.
	//
	// Enumerated values:
	//
	// JSON
	// RAW
	StorageFormat string `json:"log_storage_format"`
	// Log transfer status. ENABLE indicates that log transfer is enabled, DISABLE indicates that log transfer is disabled, and EXCEPTION indicates that log transfer is abnormal.
	//
	// Enumerated values:
	//
	// ENABLE
	// DISABLE
	// EXCEPTION
	TransferStatus string `json:"log_transfer_status"`
	// Log transfer details.
	TransferDetail TransferDetail `json:"log_transfer_detail"`
}

type TransferDetail struct {
	// Length of the transfer interval for an OBS transfer task. This parameter is required to update an OBS transfer task. The log transfer interval is specified by the combination of the values of obs_period and obs_period_unit, and must be set to one of the following: 2 min, 5 min, 30 min, 1 hour, 3 hours, 6 hours, and 12 hours.
	//
	// Enumerated values:
	//
	// 1
	// 2
	// 3
	// 5
	// 6
	// 12
	// 30
	ObsPeriod int `json:"obs_period"`
	// KMS key ID for an OBS transfer task. This parameter is required if encryption is enabled for the target OBS bucket.
	//
	// Minimum length: 36 characters
	//
	// Maximum length: 36 characters
	ObsEncryptedId string `json:"obs_encrypted_id,omitempty"`
	// Transfer file prefix of an OBS transfer task. This parameter is optional.
	//
	// The value must match the regular expression:
	//
	// ^[a-zA-Z0-9\._-]*$
	//
	// Minimum length: 1 character
	//
	// Maximum length: 64 characters
	ObsPrefixName string `json:"obs_prefix_name,omitempty"`
	// Unit of the transfer interval for an OBS transfer task. This parameter is required to update an OBS transfer task. The log transfer interval is specified by the combination of the values of obs_period and obs_period_unit, and must be set to one of the following: 2 min, 5 min, 30 min, 1 hour, 3 hours, 6 hours, and 12 hours.
	//
	// Enumerated values:
	//
	// min
	// hour
	ObsPeriodUnit string `json:"obs_period_unit"`
	// OBS bucket path, which is the log transfer destination.
	ObsTransferPath string `json:"obs_transfer_path,omitempty"`
	// OBS bucket name. This parameter is required to update an OBS transfer task.
	//
	// Minimum length: 3 characters
	//
	// Maximum length: 63 characters
	ObsBucketName string `json:"obs_bucket_name"`
	// Whether OBS bucket encryption is enabled.
	ObsEncryptedEnable bool `json:"obs_encrypted_enable,omitempty"`
	// Custom transfer path of an OBS transfer task. This parameter is optional.
	//
	// The value must match the regular expression:
	//
	// ^(/)?([a-zA-Z0-9\._-]+)(/[a-zA-Z0-9\._-]+)*(/)?$
	//
	// Minimum length: 1 character
	//
	// Maximum length: 64 characters
	ObsDirPreFixName string `json:"obs_dir_pre_fix_name,omitempty"`
	// If tag delivery is enabled, this field must contain the following host information: hostIP, hostId, hostName, pathFile, and collectTime.
	//
	// (Optional) Common fields include logStreamName, regionName, logGroupName, and projectId.
	//
	// (Optional) Enable the transfer tag: streamTag.
	Tags []string `json:"tags,omitempty"`
}

func UpdateTransfer(client *golangsdk.ServiceClient, opts UpdateTransferOpts) (*UpdateTransferResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/transfers
	raw, err := client.Put(client.ServiceURL("transfers"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateTransferResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdateTransferResponse struct {
	// Log group ID.
	//
	// Minimum length: 36 characters
	//
	// Maximum length: 36 characters
	LogGroupId string `json:"log_group_id,omitempty"`
	// Log group name.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 64 characters
	LogGroupName string `json:"log_group_name,omitempty"`
	// Log stream list.
	LogStreams []LogStreams `json:"log_streams,omitempty"`
	// Log transfer task ID.
	//
	// Minimum length: 36 characters
	//
	// Maximum length: 36 characters
	LogTransferId string `json:"log_transfer_id,omitempty"`
	// Log transfer information.
	LogTransferInfo LogTransferInfo `json:"log_transfer_info,omitempty"`
}

type LogStreams struct {
	// Log stream ID.
	//
	// Minimum length: 36 characters
	//
	// Maximum length: 36 characters
	LogStreamId string `json:"log_stream_id"`
	// Log stream name.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 64 characters
	LogStreamName string `json:"log_stream_name"`
}

type LogTransferInfo struct {
	// Information about delegated log transfer. This parameter is returned for a delegated log transfer task.
	LogAgencyTransfer LogAgencyTransfer `json:"log_agency_transfer,omitempty"`
	// Time when the log transfer task was created.
	//
	// Minimum value: 0
	//
	// Maximum value: 9999999999999
	LogCreateTime int64 `json:"log_create_time"`
	//
	// Log transfer format. The value can be RAW or JSON. RAW indicates raw log format, whereas JSON indicates JSON format. OBS transfer tasks support JSON.
	//
	// Enumerated values:
	//
	// JSON
	// RAW
	LogStorageFormat string `json:"log_storage_format"`
	// Log transfer details.
	LogTransferDetail TransferDetail `json:"log_transfer_detail"`
	//
	// Log transfer mode. cycle indicates periodical transfer, whereas realTime indicates real-time transfer. cycle is available to OBS transfer tasks and realTime is available to DIS and DMS transfer tasks.
	//
	// Enumerated values:
	//
	// cycle
	// realTime
	LogTransferMode string `json:"log_transfer_mode"`
	// Log transfer status. ENABLE indicates that log transfer is enabled, DISABLE indicates that log transfer is disabled, and EXCEPTION indicates that log transfer is abnormal.
	//
	// Enumerated values:
	//
	// ENABLE
	// DISABLE
	// EXCEPTION
	LogTransferStatus string `json:"log_transfer_status"`
	// Log transfer type. You can transfer logs to OBS.
	//
	// Enumerated values:
	//
	// OBS
	LogTransferType string `json:"log_transfer_type"`
}

type LogAgencyTransfer struct {
	// Delegator account ID.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 128 characters
	AgencyDomainId string `json:"agency_domain_id"`
	// Delegator account name.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 128 characters
	AgencyDomainName string `json:"agency_domain_name"`
	// Name of the agency created by the delegator.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 128 characters
	AgencyName string `json:"agency_name"`
	// Project ID of the delegator.
	//
	// Minimum length: 32 characters
	//
	// Maximum length: 32 characters
	AgencyProjectId string `json:"agency_project_id"`
	// Account ID of the delegated party (ID of the account that created the log transfer task).
	//
	// Minimum length: 1 character
	//
	// Maximum length: 128 characters
	BeAgencyDomainId string `json:"be_agency_domain_id"`
	//
	// Project ID of the delegated party (project ID of the account that created the log transfer task).
	//
	// Minimum length: 32 characters
	//
	// Maximum length: 32 characters
	BeAgencyProjectId string `json:"be_agency_project_id"`
}
