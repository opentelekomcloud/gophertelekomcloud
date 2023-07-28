package subscriptions

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type Subscription struct {
	RequestId       string `json:"request_id"`
	SubscriptionUrn string `json:"subscription_urn"`
}

type SubscriptionGet struct {
	TopicUrn        string `json:"topic_urn"`
	Protocol        string `json:"protocol"`
	SubscriptionUrn string `json:"subscription_urn"`
	Owner           string `json:"owner"`
	Endpoint        string `json:"endpoint"`
	Remark          string `json:"remark"`
	Status          int    `json:"status"`
}

// Extract will get the subscription object out of the commonResult object.
func (r CreateResult) Extract() (*Subscription, error) {
	s := new(Subscription)
	err := r.ExtractIntoStructPtr(s, "")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	golangsdk.Result
}

type DeleteResult struct {
	Err error
}

type ListResult struct {
	golangsdk.Result
}

func (r ListResult) Extract() ([]SubscriptionGet, error) {
	var s []SubscriptionGet
	err := r.ExtractIntoSlicePtr(&s, "subscriptions")
	if err != nil {
		return nil, err
	}
	return s, nil
}
