package connection_monitoring

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, monitorId string) (*Monitor, error) {
	raw, err := client.Get(client.ServiceURL("connection-monitors", monitorId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Monitor
	err = extract.IntoStructPtr(raw.Body, &res, "connection_monitor")
	return &res, err
}
