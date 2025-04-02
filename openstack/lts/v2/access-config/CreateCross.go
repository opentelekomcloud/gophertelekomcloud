package access_config

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateCrossOpts is a struct that contains all the parameters.
type CreateCrossOpts struct {
	// Preview of the proxy list.
	PreviewAgencyList []PreviewAgencyLogAccess `json:"preview_agency_list" required:"true"`
}

type PreviewAgencyLogAccess struct {
	// Log ingestion type.
	Type string `json:"agency_access_type" required:"true"`
	// Cross-account log ingestion configuration name.
	Name string `json:"agency_log_access" required:"true"`
	// Delegator log stream name.
	AgencyStreamName string `json:"log_agencyStream_name" required:"true"`
	// Delegator log stream ID.
	AgencyStreamId string `json:"log_agencyStream_id" required:"true"`
	// Delegator log group name.
	AgencyGroupName string `json:"log_agencyGroup_name" required:"true"`
	// Delegator log group ID.
	AgencyGroupId string `json:"log_agencyGroup_id" required:"true"`
	// Delegatee log stream name.
	StreamName string `json:"log_beAgencystream_name" required:"true"`
	// Delegatee log stream ID.
	StreamId string `json:"log_beAgencystream_id" required:"true"`
	// Delegatee log group name.
	GroupName string `json:"log_beAgencygroup_name" required:"true"`
	// Delegatee log group ID.
	GroupId string `json:"log_beAgencygroup_id" required:"true"`
	// Delegatee project ID.
	ProjectId string `json:"be_agency_project_id" required:"true"`
	// Delegator project ID.
	AgencyProjectId string `json:"agency_project_id" required:"true"`
	// Delegator account name.
	AgencyDomainName string `json:"agency_domain_name" required:"true"`
	// Agency name.
	AgencyName string `json:"agency_name" required:"true"`
}

func CrossAccess(client *golangsdk.ServiceClient, opts CreateCrossOpts) ([]AccessConfigResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2.0/{project_id}/lts/createAgencyAccess
	raw, err := client.Post(client.ServiceURL("lts", "createAgencyAccess"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res []AccessConfigResponse
	err = extract.IntoSlicePtr(raw.Body, &res, "")
	return res, err
}

type AccessConfigResponse struct {
	// Cross-account log ingestion ID.
	ID string `json:"access_config_id"`
	// Project ID.
	ProjectId string `json:"project_id"`
	// Cross-account log ingestion name.
	Name string `json:"access_config_name"`
	// Cross-account log ingestion type.
	Type string `json:"access_config_type"`
	// Log group ID.
	GroupId string `json:"group_id"`
	// Log group name.
	LogGroupName string `json:"log_group_name"`
	// Log stream ID.
	LogStreamId string `json:"log_stream_id"`
	// Log stream name.
	LogStreamName string `json:"log_stream_name"`
	// Creation time.
	CreatedAt int64 `json:"create_time"`
	// Information of the delegated ingestion.
	LogAccess *PreviewAgencyLogAccess `json:"agency_log_access"`
}
