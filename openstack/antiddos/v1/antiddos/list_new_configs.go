package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListNewConfigs(client *golangsdk.ServiceClient) (*ListConfigsResponse, error) {
	// GET /v1/{project_id}/antiddos/query_config_list
	raw, err := client.Get(client.ServiceURL("antiddos", "query_config_list"), nil, nil)
	if err != nil {
		return nil, err
	}

	var response ListConfigsResponse
	err = extract.Into(raw.Body, &response)
	return &response, err
}

type ListConfigsResponse struct {
	// List of traffic limits
	TrafficLimitedList []TrafficLimited `json:"traffic_limited_list"`
	// List of HTTP limits
	HttpLimitedList []HttpLimited `json:"http_limited_list"`
	// List of limits of numbers of connections
	ConnectionLimitedList []ConnectionLimited `json:"connection_limited_list"`
}

type TrafficLimited struct {
	// Position ID of traffic
	TrafficPosId int `json:"traffic_pos_id"`
	// Threshold of traffic per second (Mbit/s)
	TrafficPerSecond int `json:"traffic_per_second"`
	// Threshold of number of packets per second
	PacketPerSecond int `json:"packet_per_second"`
}

type HttpLimited struct {
	// Position ID of number of HTTP requests
	HttpRequestPosId int `json:"http_request_pos_id"`
	// Threshold of number of HTTP requests per second
	HttpPacketPerSecond int `json:"http_packet_per_second"`
}

type ConnectionLimited struct {
	// Position ID of access limit during cleaning
	CleaningAccessPosId int `json:"cleaning_access_pos_id"`
	// Position ID of access limit during cleaning
	NewConnectionLimited int `json:"new_connection_limited"`
	// Position ID of access limit during cleaning
	TotalConnectionLimited int `json:"total_connection_limited"`
}
