package alarm

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Query cursor. Set the value to 0 in the first query.
	// In subsequent queries, obtain the value from the response to the last request.
	Offset *int `q:"offset"`
	// Number of records on each page. The maximum value is 100.
	Limit *int `q:"limit"`
	// Specifies the name of the topic to be searched for, which is fuzzy match.
	// start with is used for the fuzzy match.
	FuzzyName string `q:"fuzzy_name"`
}

func ListTopic(client *golangsdk.ServiceClient, opts ListOpts) ([]Topics, error) {
	// GET /v2/{project_id}/lts/notifications/topics
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("lts", "notifications", "topics").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return TopicPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractTopics(pages)
}

type TopicPage struct {
	pagination.NewSinglePageBase
}

func ExtractTopics(r pagination.NewPage) ([]Topics, error) {
	var s struct {
		Topics []Topics `json:"topics"`
	}
	err := extract.Into(bytes.NewReader((r.(TopicPage)).Body), &s)
	return s.Topics, err
}

type Topics struct {
	// Topic name.
	Name string `json:"name"`
	// Specifies the resource identifier of the topic, which is unique.
	TopicUrn string `json:"topic_urn"`
	// Specifies the topic display name, which is presented as the name of the email sender in email messages.
	DisplayName string `json:"display_name"`
	// Specifies the message push policy.
	PushPolicy int `json:"push_policy"`
}
