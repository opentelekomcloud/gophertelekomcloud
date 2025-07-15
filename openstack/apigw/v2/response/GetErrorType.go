package response

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetErrorType(client *golangsdk.ServiceClient, gatewayID, groupID, id, respType string) (*ResponseInfo, error) {
	raw, err := client.Get(client.ServiceURL("apigw", "instances", gatewayID, "api-groups", groupID, "gateway-responses", id, respType),
		nil, nil)
	if err != nil {
		return nil, err
	}

	var res ResponseInfo
	err = extract.IntoStructPtr(raw.Body, &res, respType)
	return &res, err
}
