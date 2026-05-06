package cloud

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type WafSubscription struct {
	EnablePostPaidResponse
	Premium WafSubscriptionPremium `json:"premium"`
}

type WafSubscriptionPremium struct {
	Purchased bool `json:"purchased"`
	Total     int  `json:"total"`
	ELB       int  `json:"elb"`
	Dedicated int  `json:"dedicated"`
}

func Get(client *golangsdk.ServiceClient) (*WafSubscription, error) {
	raw, err := client.Get(client.ServiceURL("waf", "subscription"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res WafSubscription
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
