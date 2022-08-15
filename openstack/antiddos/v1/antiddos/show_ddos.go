package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowDDos(client *golangsdk.ServiceClient, floatingIpId string) (*ShowDDosResponse, error) {
	// GET /v1/{project_id}/antiddos/{floating_ip_id}
	raw, err := client.Get(client.ServiceURL("antiddos", floatingIpId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowDDosResponse
	err = extract.Into(raw, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type ShowDDosResponse struct {
	// Whether L7 defense has been enabled
	EnableL7 bool `json:"enable_L7,"`
	// Position ID of traffic. The value ranges from 1 to 9.
	TrafficPosId int `json:"traffic_pos_id,"`
	// Position ID of number of HTTP requests. The value ranges from 1 to 15.
	HttpRequestPosId int `json:"http_request_pos_id,"`
	// Position ID of access limit during cleaning. The value ranges from 1 to 8.
	CleaningAccessPosId int `json:"cleaning_access_pos_id,"`
	// Application type ID. Possible values: 0 1
	AppTypeId int `json:"app_type_id,"`
}
