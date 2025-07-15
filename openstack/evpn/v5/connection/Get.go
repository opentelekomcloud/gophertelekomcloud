package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, connectionId string) (*Connection, error) {
	raw, err := client.Get(client.ServiceURL("vpn-connection", connectionId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Connection
	err = extract.IntoStructPtr(raw.Body, &res, "vpn_connection")
	return &res, err
}
