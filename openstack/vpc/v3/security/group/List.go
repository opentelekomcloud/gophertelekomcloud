package group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListQueryParams struct {
	// Specifies the number of records displayed on each page.
	// The value ranges from 1 to 2000.
	Limit int `q:"limit"`
	// Specifies the start resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.
	// The value is obtained from next_marker or previous_marker in PageInfo queried last time.
	Marker string `q:"marker"`
	// Specifies the Security group ID.
	Id []string `q:"id"`
	// Specifies the Security group name.
	Name []string `q:"name"`
	// Specifies the supplementary information about the security group.
	// This field can be used to filter security groups. Multiple descriptions can be specified for filtering.
	Description []string `q:"description"`
	// Project ID
	ProjectId []string `q:"project_id"`
	// Specifies the Enterprise project ID. This field can be used to filter the security groups associated with an enterprise project.
	// The project ID can be 0 or a string that contains a maximum of 36 characters in UUID format with hyphens (-). 0 indicates the default enterprise project.
	// To obtain the security groups associated with all enterprise projects, specify all_granted_eps.
	EnterpriseProjectId []string `q:"enterprise_project_id"`
}

// This function is used to query all security groups of a tenant.
func List(client *golangsdk.ServiceClient, opts ListQueryParams) (*ListResponse, error) {
	// GET /v3/{project_id}/vpc/security-groups
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("security-groups").
		WithQueryParams(opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	// Specifies the response body for querying security groups.
	SecurityGroups []SecurityGroupListObject `json:"security_groups"`
	// Request ID
	RequestID string `json:"request_id"`
	// Specifies the pagination information.
	PageInfo PageInfo `json:"page_info"`
}

type SecurityGroupListObject struct {
	// Security group ID, which uniquely identifies the security group.
	// The value is in UUID format with hyphens (-).
	ID string `json:"id"`
	// Security group name.
	// The value can contain 1 to 64 characters, including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name"`
	// Description about the security group.
	// The value can contain up to 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description"`
	// ID of the project to which the security group belongs.
	ProjectID string `json:"project_id"`
	// Time when the security group was created.
	// The value is a UTC time in the format of yyyy-MM-ddTHH:mm:ssZ.
	CreatedAt string `json:"created_at"`
	// Time when the security group was updated.
	// The value is a UTC time in the format of yyyy-MM-ddTHH:mm:ssZ.
	UpdatedAt string `json:"updated_at"`
	// ID of the enterprise project to which the security group belongs.
	// The project ID can be 0 or a string that contains a maximum of 36 characters in UUID format with hyphens (-). 0 indicates the default enterprise project.
	EnterpriseProjectID string `json:"enterprise_project_id"`
	// Security group tags. For details, see the tag objects.
	// Value range: 0 to 20 key-value pairs.
	Tags []Tag `json:"tags"`
}

type PageInfo struct {
	// Specifies the ID of the last record in this query, which can be used in the next query.
	NextMarker string `json:"next_marker"`
	// Specifies the ID of the first record in the pagination query result.
	// When page_reverse is set to true, this parameter is used together to query resources on the previous page.
	PreviousMarker string `json:"previous_marker"`
	// Specifies the ID of the last record in the pagination query result. It is usually used to query resources on the next page. Value range: 1-200
	CurrentCount int `json:"current_count"`
}
