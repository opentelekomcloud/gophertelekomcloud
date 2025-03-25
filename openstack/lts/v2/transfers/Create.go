package transfers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Log group ID.
	// Value length: 36 characters
	LogGroupId string `json:"log_group_id" required:"true"`
	// Log stream list.
	LogStreams []LogStreams `json:"log_streams" required:"true"`
	// Log transfer information.
	LogTransferInfo *LogTransferInfo `json:"log_transfer_info" required:"true"`
}

type LogStreams struct {
	// Log stream ID.
	LogStreamId string `json:"log_stream_id" required:"true"`
	// Log stream name.
	LogStreamName string `json:"log_stream_name,omitempty"`
}

type LogTransferInfo struct {
	// Log transfer type. You can transfer logs to OBS.
	LogTransferType string `json:"log_transfer_type" required:"true"`
	// Log transfer mode. cycle indicates periodical transfer, whereas realTime indicates real-time transfer.
	// cycle is available to OBS transfer tasks and realTime is available to DIS and DMS transfer tasks.
	// Enumerated values:
	// cycle
	// realTime
	LogTransferMode string `json:"log_transfer_mode" required:"true"`
	// Log transfer format. The value can be RAW or JSON.
	// RAW indicates raw log format, whereas JSON indicates JSON format. OBS transfer tasks support JSON.
	// Enumerated values:
	// JSON
	// RAW
	LogStorageFormat string `json:"log_storage_format" required:"true"`
	// Log transfer status. ENABLE indicates that log transfer is enabled, DISABLE indicates that log transfer is disabled, and EXCEPTION indicates that log transfer is abnormal.
	// Enumerated values:
	// ENABLE
	// DISABLE
	// EXCEPTION
	LogTransferStatus string `json:"log_transfer_status" required:"true"`
	// Information about delegated log transfer.
	// This parameter is required if you transfer logs for another account.
	LogAgencyTransfer *LogAgencyTransfer `json:"log_agency_transfer,omitempty"`
	// Log transfer details.
	LogTransferDetail *TransferDetail `json:"log_transfer_detail" required:"true"`
}

type LogAgencyTransfer struct {
	// Delegator account ID.
	AgencyDomainId string `json:"agency_domain_id" required:"true"`
	// Delegator account name.
	AgencyDomainName string `json:"agency_domain_name" required:"true"`
	// Name of the agency created by the delegator.
	AgencyName string `json:"agency_name" required:"true"`
	// Project ID of the delegator.
	AgencyProjectId string `json:"agency_project_id" required:"true"`
	// Account ID of the delegated party (ID of the account that created the log transfer task).
	BeAgencyDomainId string `json:"be_agency_domain_id" required:"true"`
	// Project ID of the delegated party (project ID of the account that created the log transfer task).
	BeAgencyProjectId string `json:"be_agency_project_id" required:"true"`
}

type TransferDetail struct {
	// Length of the transfer interval for an OBS transfer task.
	// This parameter is required to create an OBS transfer task.
	// The log transfer interval is specified by the combination of the values of obs_period and obs_period_unit,
	// and must be set to one of the following: 2 min, 5 min, 30 min, 1 hour, 3 hours, 6 hours, and 12 hours.
	// Enumerated values:
	// 1
	// 2
	// 3
	// 5
	// 6
	// 12
	// 30
	ObsPeriod int `json:"obs_period" required:"true"`
	// Unit of the transfer interval for an OBS transfer task.
	// This parameter is required to create an OBS transfer task.
	// The log transfer interval is specified by the combination of the values of obs_period and obs_period_unit,
	// and must be set to one of the following: 2 min, 5 min, 30 min, 1 hour, 3 hours, 6 hours, and 12 hours.
	// Enumerated values:
	// min
	// hour
	ObsPeriodUnit string `json:"obs_period_unit" required:"true"`
	// OBS bucket name. This parameter is required to create an OBS transfer task.
	ObsBucketName string `json:"obs_bucket_name" required:"true"`
	// KMS key ID for an OBS transfer task. This parameter is required if encryption is enabled for the target OBS bucket.
	ObsEncryptedId string `json:"obs_encrypted_id,omitempty"`
	// Custom transfer path of an OBS transfer task. This parameter is optional.
	ObsDirPreFixName string `json:"obs_dir_pre_fix_name,omitempty"`
	// Transfer file prefix of an OBS transfer task. This parameter is optional.
	ObsPrefixName string `json:"obs_prefix_name,omitempty"`
	// Time zone for an OBS transfer task. For details, see Time Zone List for OBS Transfer.
	// If this parameter is specified, obs_time_zone_id must also be specified.
	ObsTimeZone string `json:"obs_time_zone,omitempty"`
	// ID of the time zone for an OBS transfer task. For details, see Time Zone List for OBS Transfer.
	// If this parameter is specified, obs_time_zone must also be specified.
	ObsTimeZoneId string `json:"obs_time_zone_id,omitempty"`
	// OBS bucket path, which is the log transfer destination.
	ObsTransferPath string `json:"obs_transfer_path,omitempty"`
	// Enterprise project ID of an OBS transfer task.
	EnterpriseProjectID string `json:"obs_eps_id,omitempty"`
	// Whether OBS bucket encryption is enabled.
	ObsEncryptedEnable bool `json:"obs_encrypted_enable,omitempty"`
	// If tag delivery is enabled, this field must contain the following host information: hostIP, hostId, hostName, pathFile, and collectTime.
	// The common fields are logStreamName, regionName, logGroupName, and projectId, which are optional.
	// The tag for enabling transfer is streamTag, which is optional.
	Tags []string `json:"tags,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*TransferResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/transfers
	raw, err := client.Post(client.ServiceURL("transfers"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res TransferResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type TransferResponse struct {
	// Log group ID.
	LogGroupId string `json:"log_group_id"`
	// Log group name.
	LogGroupName string `json:"log_group_name"`
	// Log stream list.
	LogStreams []LogStreamsResponse `json:"log_streams"`
	// Log transfer task ID.
	LogTransferId string `json:"log_transfer_id"`
	// Log transfer information.
	LogTransferInfo *LogTransferInfoResponse `json:"log_transfer_info"`
}

type LogStreamsResponse struct {
	// Log stream ID.
	LogStreamId string `json:"log_stream_id"`
	// Log stream name.
	LogStreamName string `json:"log_stream_name"`
}

type LogTransferInfoResponse struct {
	// Information about delegated log transfer.
	LogAgencyTransfer *LogAgencyTransferResponse `json:"log_agency_transfer"`
	// Time when the log transfer task was created.
	LogCreateTime int `json:"log_create_time"`
	// Log transfer format.
	LogStorageFormat string `json:"log_storage_format"`
	// Log transfer details.
	LogTransferDetail *TransferDetailResponse `json:"log_transfer_detail"`
	// Log transfer mode.
	LogTransferMode string `json:"log_transfer_mode"`
	// Log transfer status.
	LogTransferStatus string `json:"log_transfer_status"`
	// Log transfer type. You can transfer logs to OBS.
	LogTransferType string `json:"log_transfer_type"`
}

type LogAgencyTransferResponse struct {
	// Delegator account ID.
	AgencyDomainId string `json:"agency_domain_id"`
	// Delegator account name.
	AgencyDomainName string `json:"agency_domain_name"`
	// Name of the agency created by the delegator.
	AgencyName string `json:"agency_name"`
	// Project ID of the delegator.
	AgencyProjectId string `json:"agency_project_id"`
	// Account ID of the delegated party (ID of the account that created the log transfer task).
	BeAgencyDomainId string `json:"be_agency_domain_id"`
	// Project ID of the delegated party (project ID of the account that created the log transfer task).
	BeAgencyProjectId string `json:"be_agency_project_id"`
}

type TransferDetailResponse struct {
	// Length of the transfer interval for an OBS transfer task.
	ObsPeriod int `json:"obs_period"`
	// Unit of the transfer interval for an OBS transfer task.
	ObsPeriodUnit string `json:"obs_period_unit"`
	// OBS bucket name.
	ObsBucketName string `json:"obs_bucket_name"`
	// KMS key ID for an OBS transfer task.
	ObsEncryptedId string `json:"obs_encrypted_id"`
	// Custom transfer path of an OBS transfer task.
	ObsDirPreFixName string `json:"obs_dir_pre_fix_name"`
	// Transfer file prefix of an OBS transfer task.
	ObsPrefixName string `json:"obs_prefix_name"`
	// Time zone for an OBS transfer task.
	ObsTimeZone string `json:"obs_time_zone"`
	// ID of the time zone for an OBS transfer task.
	ObsTimeZoneId string `json:"obs_time_zone_id"`
	// OBS bucket path, which is the log transfer destination.
	ObsTransferPath string `json:"obs_transfer_path"`
	// Enterprise project ID of an OBS transfer task.
	EnterpriseProjectID string `json:"obs_eps_id"`
	// Whether OBS bucket encryption is enabled.
	ObsEncryptedEnable bool `json:"obs_encrypted_enable"`
	// If tag delivery is enabled.
	Tags []string `json:"tags,omitempty"`
}
