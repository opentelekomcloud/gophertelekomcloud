package cluster

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type DeleteOpts struct {
	// KeepBackup Number of backup log files. Retain the default value 0.
	KeepBackup int `json:"keep_last_manual_backup"`
}

// Delete is used to delete a cluster.
// Send request DELETE /v1.1/{project_id}/clusters/{cluster_id}
func Delete(client *golangsdk.ServiceClient, id string, reqOpts *DeleteOpts) (*JobId, error) {
	b, err := build.RequestBody(reqOpts, "")
	if err != nil {
		return nil, err
	}

	r, err := client.DeleteWithBody(client.ServiceURL(clustersEndpoint, id), b, nil)
	if err != nil {
		return nil, err
	}

	var resp JobId
	err = extract.Into(r.Body, &resp)

	return &resp, err
}
