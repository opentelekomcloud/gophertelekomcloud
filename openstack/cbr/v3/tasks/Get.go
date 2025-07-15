package tasks

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*OperationLog, error) {
	raw, err := client.Get(client.ServiceURL("operation-logs", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res OperationLog
	err = extract.IntoStructPtr(raw.Body, &res, "operation_log")
	return &res, err
}
