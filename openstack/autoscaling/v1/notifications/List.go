package notifications

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// List returns the notification configurations for an AS group.
func List(client *golangsdk.ServiceClient, scalingGroupID string) ([]Notification, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("scaling_notification", scalingGroupID).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return NotificationPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractNotifications(pages)
}

type NotificationPage struct {
	pagination.NewSinglePageBase
}

func ExtractNotifications(page pagination.NewPage) ([]Notification, error) {
	var res struct {
		Topics []Notification `json:"topics"`
	}
	err := extract.Into(bytes.NewReader(page.(NotificationPage).Body), &res)
	return res.Topics, err
}
