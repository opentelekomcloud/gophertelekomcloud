package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowDDosStatus(client *golangsdk.ServiceClient, floatingIpId string) (string, error) {
	// GET /v1/{project_id}/antiddos/{floating_ip_id}/status
	raw, err := client.Get(client.ServiceURL("antiddos", floatingIpId, "status"), nil, nil)
	if err != nil {
		return "", err
	}

	var res struct {
		Status string `json:"status"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Status, err
}
