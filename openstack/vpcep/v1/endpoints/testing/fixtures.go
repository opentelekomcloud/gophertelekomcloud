package testing

import (
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/endpoints"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/vpcep/v1/services"
)

const (
	createRequest = `
{
  "subnet_id": "68bfbcc1-dff2-47e4-a9d4-332b9bc1b8de",
  "vpc_id": "84758cf5-9c62-43ae-a778-3dbd8370c0a4",
  "tags": [
    {
      "key": "test1",
      "value": "test1"
    }
  ],
  "endpoint_service_id": "e0c748b7-d982-47df-ba06-b9c8c7650c1a",
  "enable_dns": true
}
`
	createResponse = `
{
  "id": "4189d3c2-8882-4871-a3c2-d380272eed83",
  "service_type": "interface",
  "marker_id": 322312312312,
  "status": "creating",
  "vpc_id": "84758cf5-9c62-43ae-a778-3dbd8370c0a4",
  "enable_dns": true,
  "endpoint_service_name": "test123",
  "endpoint_service_id": "e0c748b7-d982-47df-ba06-b9c8c7650c1a",
  "project_id": "6e9dfd51d1124e8d8498dce894923a0d",
  "created_at": "2018-01-30T07:42:01.174",
  "updated_at": "2018-01-30T07:42:01.174",
  "tags": [
    {
      "key": "test1",
      "value": "test1"
    }
  ]
}
`

	listResponse = `
{
  "endpoints": [
    {
      "id": "03184a04-95d5-4555-86c4-e767a371ff99",
      "status": "accepted",
      "ip": "192.168.0.232",
      "marker_id": 16777337,
      "active_status": [
		"active"
	  ],
      "vpc_id": "84758cf5-9c62-43ae-a778-3dbd8370c0a4",
      "service_type": "interface",
      "project_id": "295dacf46a4842fcbf7844dc2dc2489d",
      "subnet_id": "68bfbcc1-dff2-47e4-a9d4-332b9bc1b8de",
      "enable_dns": true,
      "dns_name": "test123",
      "created_at": "2018-10-18T06:49:46Z",
      "updated_at": "2018-10-18T06:49:50Z",
      "endpoint_service_id": "5133655d-0e28-4090-b669-13f87b355c78",
      "endpoint_service_name": "test123",
      "whitelist": [
        "127.0.0.1"
      ],
      "enable_whitelist": true,
      "tags": [
        {
          "key": "test1",
          "value": "test1"
        }
      ]
    },
    {
      "id": "43b0e3b0-eec9-49da-866b-6687b75f9fe5",
      "status": "accepted",
      "ip": "192.168.0.115",
      "marker_id": 16777322,
      "active_status": [
		"active"
	  ],
      "vpc_id": "84758cf5-9c62-43ae-a778-3dbd8370c0a4",
      "service_type": "interface",
      "project_id": "295dacf46a4842fcbf7844dc2dc2489d",
      "subnet_id": "65528a22-59a1-4972-ba64-88984b3207cd",
      "enable_dns": true,
      "dns_name": "test123",
      "created_at": "2018-10-18T06:36:20Z",
      "updated_at": "2018-10-18T06:36:24Z",
      "endpoint_service_id": "5133655d-0e28-4090-b669-13f87b355c78",
      "endpoint_service_name": "test123",
      "whitelist": [
        "127.0.0.1"
      ],
      "enable_whitelist": true,
      "tags": [
        {
          "key": "test1",
          "value": "test1"
        }
      ]
    }
  ],
  "total_count": 17
}
`
)

var expected = &endpoints.Endpoint{
	ID:          "4189d3c2-8882-4871-a3c2-d380272eed83",
	ServiceType: services.ServiceTypeInterface,
	MarkerID:    322312312312,
	Status:      endpoints.StatusCreating,
	VpcID:       "84758cf5-9c62-43ae-a778-3dbd8370c0a4",
	EnableDNS:   true,
	ServiceName: "test123",
	ServiceID:   "e0c748b7-d982-47df-ba06-b9c8c7650c1a",
	ProjectID:   "6e9dfd51d1124e8d8498dce894923a0d",
	CreatedAt:   "2018-01-30T07:42:01.174",
	UpdatedAt:   "2018-01-30T07:42:01.174",
	Tags:        []tags.ResourceTag{{Key: "test1", Value: "test1"}},
}
