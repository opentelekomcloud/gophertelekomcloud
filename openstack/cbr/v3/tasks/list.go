package tasks

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type ListOpts struct {
	EndTime             string `q:"end_time"`
	EnterpriseProjectId string `q:"enterprise_project_id"`
	Limit               int    `q:"limit"`
	Offset              int    `q:"offset"`
	OperationType       string `q:"operation_type"`
	ProviderId          string `q:"provider_id"`
	ResourceId          string `q:"resource_id"`
	ResourceName        string `q:"resource_name"`
	StartTime           string `q:"start_time"`
	Status              string `q:"status"`
	VaultId             string `q:"vault_id"`
	VaultName           string `q:"vault_name"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]OperationLog, error) {
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("operation-logs")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []OperationLog
	return res, extract.IntoSlicePtr(raw.Body, &res, "operation_log")
}
